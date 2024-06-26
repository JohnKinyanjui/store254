package mpesa_plugin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Mpesa struct {
	//if set to true the access roken will be reused
	//until it expires
	//otherwise each request will always get a new token
	//which will slow down your requests
	//the default is true
	//For sandbox use false and for production use true
	Live                bool
	ConsumerKey         string
	ConsumerSecret      string
	DefaultTimeOut      time.Duration
	DefaultPassKey      string
	DefaultC2BShortCode string

	cache map[string]AccessTokenResponse
}

func New(ConsumerKey, ConsumerSecret string, live bool) Mpesa {
	TokenCache := make(map[string]AccessTokenResponse)
	return Mpesa{
		Live:           live,
		ConsumerKey:    ConsumerKey,
		ConsumerSecret: ConsumerSecret,
		DefaultTimeOut: 20 * time.Second,
		cache:          TokenCache,
	}
}

// SetDefaultPassKey You can set the default pass key
// Over here so that you dont have to pass it each time
// You are sending an StkRequest
func (m *Mpesa) SetDefaultPassKey(passKey string) {
	m.DefaultPassKey = passKey

}

// SetDefaultTimeOut this will set the connection timeout to Mpesa
// the default is 20 seconds when sending an http request
func (m *Mpesa) SetDefaultTimeOut(timeOut time.Duration) {
	m.DefaultTimeOut = timeOut
}

// SetLiveMode  changes from production to sandbox and viceversa
// at runtime.
func (m *Mpesa) SetLiveMode(mode bool) {
	m.Live = mode
}

// StkPushRequest send an Mpesa express request
// note if you have already set a DefaultPassKey you don't have to pass
// a pass key here its optional
// If you also set DefaultC2BShortCode you dont have to pass BusinessShortCode to the StKPushRequestBody
// the default will be used if you don't pass it
func (m *Mpesa) StkPushRequest(body StKPushRequestBody, passKey ...string) (*StkPushResult, error) {
	var stkPassKey string
	if len(passKey) > 0 {
		stkPassKey = passKey[0]
	} else {
		stkPassKey = m.DefaultPassKey
	}
	if IsEmpty(stkPassKey) {
		return nil, errors.New("pass key is needed set a default pass key or pass it in ths function")
	}
	if IsEmpty(body.BusinessShortCode) {
		body.BusinessShortCode = m.DefaultC2BShortCode
	}
	err := body.Validate()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	fTime := t.Format("20060102150405")
	requestBody := StkRequestFullBody{
		StKPushRequestBody: body,
		Password:           GeneratePassword(body.BusinessShortCode, stkPassKey, fTime),
		Timestamp:          fTime,
		TransactionType:    CustomerPayBillOnline,
		PartyA:             body.PhoneNumber,
		PartyB:             body.PartyB,
	}
	var stkPushResult StkPushResult
	err = m.sendAndProcessStkPushRequest(m.getStkPush(), requestBody, &stkPushResult, nil)
	return &stkPushResult, err
}

// StkPushVerification use this to confirm your stk push if it was a failure or success
// CheckoutRequestID is the CheckoutRequestID you got when you sent the StkPushRequest request
// you dont have to send a passkey if you have a DefaultPassKey set
func (m *Mpesa) StkPushVerification(CheckoutRequestID string, BusinessShortCode string, passKey ...string) (*StkPushQueryResponseBody, error) {
	var stkPassKey string
	if len(passKey) > 0 {
		stkPassKey = passKey[0]
	} else {
		stkPassKey = m.DefaultPassKey
	}
	if IsEmpty(stkPassKey) {
		return nil, errors.New("pass key is needed set a default pass key or pass it in ths function")
	}
	t := time.Now()
	fTime := t.Format("20060102150405")
	body := StkPushQueryRequestBody{
		BusinessShortCode: BusinessShortCode,
		Password:          GeneratePassword(BusinessShortCode, stkPassKey, fTime),
		Timestamp:         fTime,
		CheckoutRequestID: CheckoutRequestID,
	}
	var stkPushResult StkPushQueryResponseBody
	err := m.sendAndProcessStkPushRequest(m.getStkPushQuery(), body, &stkPushResult, nil)
	return &stkPushResult, err
}

func (m *Mpesa) StkPushQuery(body StkPushQueryRequestBody, passKey ...string) (*StkPushQueryResponseBody, error) {
	var stkPassKey string
	if len(passKey) > 0 {
		stkPassKey = passKey[0]
	} else {
		stkPassKey = m.DefaultPassKey
	}
	if IsEmpty(stkPassKey) {
		return nil, errors.New("pass key is needed set a default pass key or pass it in ths function")
	}
	if body.Timestamp == "" {
		t := time.Now()
		fTime := t.Format("20060102150405")
		body.Timestamp = fTime
		body.Password = GeneratePassword(body.BusinessShortCode, stkPassKey, fTime)
	}

	var stkPushResult StkPushQueryResponseBody
	err := m.sendAndProcessStkPushRequest(m.getStkPushQuery(), body, &stkPushResult, nil)
	return &stkPushResult, err
}

func (m *Mpesa) sendAndProcessStkPushRequest(url string, data interface{}, respItem interface{}, extraHeader map[string]string) error {
	if reflect.ValueOf(respItem).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return errors.New("response should be a pointer")
	}
	token, err := m.GetAccessToken()
	if err != nil {

		return err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	///auth := "Bearer " + token.AccessToken
	//headers["authorization"]= auth
	headers["Authorization"] = "Bearer " + token.AccessToken
	///add the extra headers
	//Get union of the headers
	for k, v := range extraHeader {
		headers[k] = v
	}
	resp, err := postRequest(url, data, headers, m.DefaultTimeOut)
	if err != nil {

		return err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)

		return &RequestError{Message: string(b), StatusCode: resp.StatusCode}
	}

	///var respe map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(respItem); err != nil {

		return errors.New("error converting from json")
	}

	return nil

}

// GetAccessToken will get the token to be used to query data
func (m *Mpesa) GetAccessToken() (*AccessTokenResponse, error) {
	if m.cache == nil {
		m.cache = make(map[string]AccessTokenResponse)
	}
	//check if we have a cached Token and return it instead
	//if we have allowed caching

	req, err := http.NewRequest(http.MethodGet, m.getAccessTokenUrl(), nil)
	if err != nil {
		return nil, err
	}
	log.Println(req.URL.String())
	req.SetBasicAuth(m.ConsumerKey, m.ConsumerSecret)

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{
		Timeout: m.DefaultTimeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		b, _ := ioutil.ReadAll(resp.Body)
		if string(b) == "" {
			return nil, &RequestError{Message: "Error getting token", StatusCode: resp.StatusCode}

		}
		return nil, &RequestError{Message: string(b), StatusCode: resp.StatusCode}
	}
	var token AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {

		return nil, errors.New("error converting from json")
	}

	if m.cache == nil {
		m.cache = make(map[string]AccessTokenResponse)

	}

	return &token, nil
}

func postRequest(url string, data interface{}, headers map[string]string, timeOut time.Duration) (*http.Response, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	log.Println(string(b))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: timeOut,
	}
	return client.Do(req)

}

type RequestError struct {
	StatusCode int

	Message string
	Url     string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("url is: %s \n status code is: %d \n  and body is : %s", r.Url, r.StatusCode, r.Message)

}

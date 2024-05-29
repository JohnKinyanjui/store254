package http

import "errors"

type Http struct {
	url string
	err error
}

func HttpClient(url string) *Http {
	if len(url) == 0 {
		return &Http{err: errors.New("url is missing")}
	}

	return &Http{url: url}
}

func (http *Http) Headers(headers *map[string]interface{}) *Http {
	return http
}

func (http *Http) Post(body *map[string]interface{}) map[string]interface{} {
	var res map[string]interface{}

	return res
}

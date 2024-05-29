package mpesa_plugin

const MpesaLiveUrl = "https://api.safaricom.co.ke/"
const MpesaSandboxUrl = "https://sandbox.safaricom.co.ke/"
const tokenUrl = "oauth/v1/generate?grant_type=client_credentials"
const stkPush = "mpesa/stkpush/v1/processrequest"
const stkPushQuery = "mpesa/stkpushquery/v1/query"

func (m *Mpesa) getStkPush() string {

	return m.getBaseUrl() + stkPush

}
func (m *Mpesa) getStkPushQuery() string {

	return m.getBaseUrl() + stkPushQuery

}
func (m *Mpesa) getAccessTokenUrl() string {

	return m.getBaseUrl() + tokenUrl
}

func (m *Mpesa) getBaseUrl() string {
	if !m.Live {

		return MpesaSandboxUrl
	}

	return MpesaLiveUrl

}

package mpesa_plugin

import (
	"encoding/base64"
	"fmt"
)

const (
	SandBox                = "sandbox"
	TransactionReversal    = "TransactionReversal"
	AccountBalance         = "AccountBalance"
	CustomerPayBillOnline  = "CustomerPayBillOnline"
	TransactionStatusQuery = "TransactionStatusQuery"

	// Identifier Types
	// see https://developer.safaricom.co.ke/docs#identifier-types
	PayBillIdentifier    = "4"
	TillNumberIdentifier = "2"
	MSISDNIdentifier     = "1"
	AccessToken          = "access_token"
)

func GeneratePassword(shortCode, passkey, time string) string {
	password := fmt.Sprintf("%s%s%s", shortCode, passkey, time)
	return base64.StdEncoding.EncodeToString([]byte(password))

}

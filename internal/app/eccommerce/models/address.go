package models

import (
	"errors"
)

type Address struct {
	Name        string                 `json:"name" `
	PhoneNumber string                 `json:"phone_number"`
	Email       string                 `json:"email" pg:"email"`
	Street      string                 `json:"street" bun:"street"`
	City        string                 `json:"city" bun:"city"`
	State       string                 `json:"state" bun:"state"`
	PostalCode  string                 `json:"postal_code" bun:"postal_code"`
	Country     string                 `json:"country" bun:"country"`
	Metadata    map[string]interface{} `json:"metadata"`
}

func (addr *Address) Validation() error {
	fields := map[string]string{
		"Name":         addr.Name,
		"Phone Number": addr.PhoneNumber,
		"Email":        addr.Email,
		"Street":       addr.Street,
		"City":         addr.City,
		"State":        addr.State,
		"Country":      addr.Country,
		"PostalCode":   addr.PostalCode,
	}

	for name, value := range fields {
		if name != "Phone Number" && name != "PostalCode" {
			if len(value) == 0 {
				return errors.New(name + " is missing in the address provided")
			}
		}
	}

	return nil
}

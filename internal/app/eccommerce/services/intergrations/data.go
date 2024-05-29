package intergration_service

import "github.com/google/uuid"

var availableOauths = []string{"google"}

type IntegrationOauthPostData struct {
	StoreId      uuid.UUID         `json:"store_id"`
	Intergration string            `json:"intergration"`
	Credentials  map[string]string `json:"credentials"`
}

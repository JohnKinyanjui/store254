package authconfigs_service

import (
	"github.com/google/uuid"
)

type PostGoogleData struct {
	ID           uuid.UUID `json:"id"`
	StoreID      uuid.UUID `json:"store_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	RedirectURI  string    `json:"redirect_uri"`
}

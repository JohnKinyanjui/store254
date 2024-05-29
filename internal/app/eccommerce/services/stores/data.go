package store_service

import "github.com/google/uuid"

type StoreSearchParam struct {
	StoreId      string
	UserId       string
	Name         string
	Collection   []string
	MinimumPrice string
	MaximumPrice string
	Rating       string
	UserType     string
}

type StorePostData struct {
	StoreId     uuid.UUID `json:"store_id"`
	UserId      uuid.UUID `json:"user_id"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	ContactInfo struct {
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Facebook    string `json:"facebook"`
		Twitter     string `json:"twitter"`
		Instagram   string `json:"instagram"`
	} `json:"contact_info"`
}

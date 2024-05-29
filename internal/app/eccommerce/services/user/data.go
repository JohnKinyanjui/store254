package user_service

type UserWaitlistData struct {
	Email    string `json:"email"`
	Approved bool   `json:"approved"`
}

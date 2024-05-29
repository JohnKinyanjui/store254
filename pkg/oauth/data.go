package oauth

import "time"

var (
	outhStates = make(map[string]time.Time, 0)
)

type OauthUser struct {
	UserId               string
	Uid                  string
	Picture              string
	UserName             string
	Email                string
	Token                string
	GithubInstallationId string
}

type googleUser struct {
	Sub           string `json:"sub"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type gitHubUser struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type Email struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
	// "verified": true,
	// "visibility": "private"
}

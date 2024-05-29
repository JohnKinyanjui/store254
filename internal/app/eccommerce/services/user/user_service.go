package user_service

import "eccomerce_apis/internal/config"

func CreateUserWaitlist(cfg *config.Config, data *UserWaitlistData) error {
	return createWaitlist(cfg, data)
}

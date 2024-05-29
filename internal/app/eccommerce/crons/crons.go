package eccommerce_crons

import (
	"eccomerce_apis/internal/config"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func StartAllCrons(config *config.Config) {
	log.Println("[cron jobs started running]")

	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(5).Hours().Tag("store_customer_subscription").Do(func() {
		CheckSubscription(config)
	})

	if err != nil {
		log.Println("[cron jobs stopped]")
		return
	}

	s.StartAsync()
}

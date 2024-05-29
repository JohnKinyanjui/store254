package security

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type StoreUserSession struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	ExpiryAt time.Time `json:"expiry_at"`
}

func (s *StoreUserSession) Params(sessionId string) {
	id, err := uuid.Parse(sessionId)
	if err != nil {
		return
	}

	s.Id = id
}

func (s *StoreUserSession) Create(db *pg.DB) error {

	return nil
}

func (s *StoreUserSession) Delete() {}

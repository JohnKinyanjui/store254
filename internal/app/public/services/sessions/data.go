package sessions_service

import "github.com/google/uuid"

var sessions = make(map[string]Session, 0)

type Session struct {
	SessionId  uuid.UUID
	StoreId    uuid.UUID
	CustomerId uuid.UUID
	Role       string
}

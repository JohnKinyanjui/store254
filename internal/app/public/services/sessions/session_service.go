package sessions_service

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func GenerateUserSession(rd *redis.Client, session *Session) (*Session, error) {
	return generateUserSession(rd, session)
}

func GetUserSession(rd *redis.Client, sessionId uuid.UUID) (*Session, error) {
	return getUserSession(rd, sessionId.String())
}

package sessions_service

import (
	"context"
	"eccomerce_apis/pkg/middlewares"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func generateUserSession(rd *redis.Client, session *Session) (*Session, error) {
	sessionData := map[string]string{
		"store_id":    session.StoreId.String(),
		"customer_id": session.CustomerId.String(),
		"expiry_at":   time.Duration(730 * time.Hour).String(),
	}

	ecryptedData, err := middlewares.CredEncrypt().Encrypt(sessionData)
	if err != nil {
		return nil, err
	}

	sessions[session.SessionId.String()] = *session
	rd.Set(context.Background(), session.SessionId.String(), ecryptedData, time.Duration(730*time.Hour))

	// CANNOT LOGIN TWO TIMES Í
	userSessions := rd.Get(context.Background(), session.CustomerId.String())
	if len(userSessions.String()) > 0 {
		sessionsFound := strings.Split(userSessions.String(), ",")
		rd.Del(context.Background(), sessionsFound...)
	} else {
		rd.Set(context.Background(), session.CustomerId.String(), session.SessionId.String(), 0)
	}

	return session, nil
}

func getUserSession(rd *redis.Client, sessionId string) (*Session, error) {
	session, ok := sessions[sessionId]
	if ok {
		return &session, nil
	}

	ses, err := rd.Get(context.Background(), sessionId).Result()
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to get session from redis").Log()
	}

	ecryptedData, err := middlewares.CredEncrypt().Decrypt(ses)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to get session").Log()
	}

	customerId, _ := uuid.Parse(ecryptedData["customer_id"])
	storeId, _ := uuid.Parse(ecryptedData["store_id"])

	session = Session{
		StoreId:    storeId,
		CustomerId: customerId,
	}

	sessions[sessionId] = session

	return &session, nil
}

func refreshUserSession() {}

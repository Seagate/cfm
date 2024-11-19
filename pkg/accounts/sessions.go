// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package accounts

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

const (
	defaultSessionSeconds float64 = 60 * 30.0
)

type SessionInformation struct {
	Id       string
	Token    string
	Created  time.Time
	Updated  time.Time
	Username string
	Timeout  float64
}

// map of session id string to session information
var sessions map[string]*SessionInformation

func init() {
	sessions = make(map[string]*SessionInformation)
}

var (
	ErrInvalidCredentials = errors.New("invalid user credentials")
	ErrGenerateSessionId  = errors.New("session id creation failure")
)

// CreateSession: create a new session and return it's information
func CreateSession(ctx context.Context, username, password string) (*SessionInformation, error) {
	logger := klog.FromContext(ctx)

	var session *SessionInformation

	// Validate user credentials
	valid, err := AccountsHandler().ValidAccount(username, password)
	if !valid || err != nil {
		logger.V(1).Error(err, "failure: create session", "username", username)
		return nil, ErrInvalidCredentials
	}

	// Example uuid: ee0328d9-258a-4e81-976e-b75aa4a2d8f5
	token := uuid.New().String()
	token = strings.ReplaceAll(token, "-", "")

	// Generate a new session id
	id, err := GenerateSessionId()
	if err != nil {
		logger.V(1).Error(err, "failure: generate session id", "username", username)
		return nil, ErrGenerateSessionId
	}

	// Store the session information using the session id
	created := time.Now()
	session = &SessionInformation{Id: id, Token: token, Created: created, Updated: created, Username: username, Timeout: defaultSessionSeconds}
	sessions[id] = session

	logger.V(1).Info("success: created session:", "session", session)

	return session, nil
}

// GetSessions: retrieve the map of all sessions
func GetSessions() map[string]*SessionInformation {
	return sessions
}

// UpdateSessionTime: update the
func UpdateSessionTime(id string) error {
	value, ok := sessions[id]
	if ok {
		value.Updated = time.Now()
		return nil
	}
	return fmt.Errorf("(%s) was not found in the table of sessions", id)
}

// IsSessionTokenActive: return true if session is active
func IsSessionTokenActive(ctx context.Context, token string) bool {
	logger := klog.FromContext(ctx)
	active := false

	for id, session := range sessions {
		if session.Token == token {
			// Make sure that the retrieved session has not expired
			current := time.Now()
			duration := current.Sub(session.Updated)
			if duration.Seconds() > session.Timeout {
				// Session has expired, remove it
				logger.V(1).Info("session expired", "id", id, "current", current, "updated", session.Updated, "duration", duration.Seconds())
				delete(sessions, id)
			} else {
				// Session is still active and being used, so update time used
				active = true
				session.Updated = time.Now()
			}
			break
		}
	}

	return active
}

// IsSessionIdActive: return true if session is active
func IsSessionIdActive(ctx context.Context, id string) bool {
	logger := klog.FromContext(ctx)
	active := false
	session, ok := sessions[id]

	if !ok || session == nil {
		logger.V(1).Info("session was not found in lookup table", "id", id, "ok", ok, "session", session)
	}

	if ok && session != nil {
		// Make sure that the retrieved session has not expired
		current := time.Now()
		duration := current.Sub(session.Updated)
		if duration.Seconds() > session.Timeout {
			// Session has expired, remove it
			logger.V(1).Info("session expired", "current", current, "updated", session.Updated, "duration", duration.Seconds())
			delete(sessions, id)
		} else {
			// Session is still active and being used, so update time used
			active = true
			session.Updated = time.Now()
		}
	}

	return active
}

// GetSessionInformation: return session information
func GetSessionInformation(ctx context.Context, id string) *SessionInformation {
	info, ok := sessions[id]
	if !ok {
		return nil
	} else {
		return info
	}
}

// DeleteSession: delete the session and return session information
func DeleteSession(ctx context.Context, id string) *SessionInformation {
	info, ok := sessions[id]
	if !ok {
		return nil
	} else {
		delete(sessions, id)
		return info
	}
}

const randomCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

func randomString(r *rand.Rand, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomCharset[r.Intn(len(randomCharset))]
	}
	return string(b)
}

// generateUniqueKey - Generic function for generating a unique key for an existing map
func generateUniqueKey(r *rand.Rand, length int, existingMap map[string]interface{}, maxAttempts int) (string, error) {
	for attempts := 0; attempts < maxAttempts; attempts++ {
		key := randomString(r, length)
		if _, exists := existingMap[key]; !exists {
			return key, nil
		}
	}
	return "", fmt.Errorf("failed to generate a unique key after maximum attempts")
}

// GenerateSessionId - Generates a new session id, consisting of 10 random alpha-numeric chars
func GenerateSessionId() (string, error) {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	maxAttempts := 100
	keyLength := 10

	// Convert sessions map to map[string]interface{}
	interfaceMap := make(map[string]interface{})
	for k, v := range sessions {
		interfaceMap[k] = v
	}

	// Example of generating and storing unique keys
	id, err := generateUniqueKey(r, keyLength, interfaceMap, maxAttempts)
	if err != nil {
		return "", fmt.Errorf("failure: session id generation: %s", err)
	}

	return id, nil
}

// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package accounts

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

const (
	defaultSessionSeconds float64 = 60 * 30.0
	maxSessionId          int     = math.MaxInt
)

// Track the last session id used, increment until a max is reached, then reset to 1
var activeSessionId int = 0

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

// CreateSession: create a new session and return it's information
func CreateSession(ctx context.Context, username, password string) *SessionInformation {

	var session *SessionInformation

	// Validate user credentials
	valid, err := AccountsHandler().ValidAccount(username, password)
	if valid && err == nil {
		session = CreateSessionToken(ctx, username)
	}

	return session
}

// CreateSessionToken: create a token and add it to the map, return a new session id
func CreateSessionToken(ctx context.Context, user string) *SessionInformation {
	logger := klog.FromContext(ctx)

	// Example uuid: ee0328d9-258a-4e81-976e-b75aa4a2d8f5
	token := uuid.New().String()
	token = strings.ReplaceAll(token, "-", "")

	// Determine the next session id to use
	activeSessionId += 1
	if activeSessionId > maxSessionId {
		logger.V(1).Info("session id was reset", "maxSessionId", maxSessionId)
		activeSessionId = 1
	}
	id := strconv.Itoa(activeSessionId)

	// Store the session information using the session id (an integer from 1..max)
	created := time.Now()
	session := &SessionInformation{Id: id, Token: token, Created: created, Updated: created, Username: user, Timeout: defaultSessionSeconds}
	sessions[id] = session

	return session
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

package session

import (
	"http/web"
	"huntsub/huntsub-map-server/o/auth/session"
)

const (
	errReadSessonFailed   = web.InternalServerError("read session failed")
	errSessionNotFound    = web.Unauthorized("session not found")
	errUnauthorizedAccess = web.Unauthorized("unauthorized access")
)

func Get(sessionID string) (*session.Session, error) {
	var s, err = session.GetByID(sessionID)
	if err != nil {
		sessionLog.Error(err)
		return nil, errReadSessonFailed
	}
	if s == nil {
		return nil, errSessionNotFound
	}
	return s, nil
}

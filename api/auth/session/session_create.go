package session

import (
	"http/web"
	"huntsub/huntsub-map-server/o/auth/session"
	"huntsub/huntsub-map-server/o/org/user"
	"time"
)

func New(u *user.User) (*session.Session, error) {
	var s = &session.Session{
		UserID:   u.ID,
		Username: u.Username,
		CTime:    time.Now().Unix(),
	}

	err := s.Create()
	if err != nil {
		sessionLog.Error(err)
		return nil, web.InternalServerError("save session failed")
	}
	return s, nil
}

func MustNew(u *user.User) *session.Session {
	s, e := New(u)
	if e != nil {
		panic(e)
	}
	return s
}

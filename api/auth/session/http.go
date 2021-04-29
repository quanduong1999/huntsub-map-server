package session

import (
	"huntsub/huntsub-map-server/o/auth/session"
	"huntsub/huntsub-map-server/x/mlog"
	"net/http"
	"strings"
)

var accessToken = "token"
var authorization = "Authorization"
var sessionLog = mlog.NewTagLog("session_log")

func MustGet(r *http.Request) *session.Session {
	var sessionID = r.URL.Query().Get(accessToken)
	var s, e = Get(sessionID)
	if e != nil {
		panic(e)
	}
	return s
}

func MustAuthScope(r *http.Request) *session.Session {
	var sess = r.Header.Get(authorization)
	var str = strings.Split(sess, " ")
	var s, e = Get(str[1])
	if e != nil {
		panic(e)
	}
	return s
}

func MustClear(r *http.Request) {
	var sessionID = r.URL.Query().Get(accessToken)
	var e = session.MarkDelete(sessionID)
	if e != nil {
		sessionLog.Error(e, "remove session")
	}
}

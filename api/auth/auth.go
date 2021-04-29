package auth

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/user"
	"net/http"
	"strings"
)

var obj = event.ObjectEventSource

type AuthServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewAuthServer() *AuthServer {
	var s = &AuthServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/login", s.HandleLogin)
	s.HandleFunc("/me", s.handleMe)
	s.HandleFunc("/my_settings", s.handleMySettings)
	s.HandleFunc("/logout", s.HandleLogout)
	s.HandleFunc("/change_pass", s.handleChangePass)
	return s
}

func (s *AuthServer) handleMe(w http.ResponseWriter, r *http.Request) {
	s.SendJson(w, session.MustAuthScope(r))
}

func (s *AuthServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		Username string
		Password string
		Scope    string
		Auto     bool
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}

	if err = u.ComparePassword(body.Password); err != nil {
		s.SendError(w, err)
		return
	}

	if !u.Verify {
		s.SendError(w, errUserNotVerify)
		return
	}

	var ses = session.MustNew(u)

	w.Header().Set("Huntsub-Token", ses.ID)
	s.SendData(w, map[string]interface{}{
		"user":    u,
		"session": ses,
	})
}

func (s *AuthServer) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session.MustClear(r)
	s.SendData(w, nil)
}

func (s *AuthServer) handleChangePass(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		OldPass   string `json:"old_pass"`
		NewPass   string `json:"new_pass"`
		ReNewPass string `json:"re_new_pass"`
		Username  string `json:"username"`
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}

	if err := u.ComparePassword(body.OldPass); err != nil {
		s.SendError(w, err)
		return
	}
	u.UpdatePass(body.NewPass)

	s.SendData(w, map[string]interface{}{
		"status": "success",
	})
	obj.EmitUpdate(u)
}

package auth

import (
	"huntsub/huntsub-map-server/api/auth/session"
	"huntsub/huntsub-map-server/o/org/user"
	"net/http"
)

func (s *AuthServer) handleMySettings(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	me, _ := user.GetByID(u.UserID)
	s.SendData(w, map[string]interface{}{
		"me": me,
	})
}

package api

import (
	"http/web"
	"huntsub/huntsub-map-server/api/action/follow"
	"huntsub/huntsub-map-server/api/action/like"
	"huntsub/huntsub-map-server/api/action/report"
	"huntsub/huntsub-map-server/api/action/share"
	"huntsub/huntsub-map-server/api/auth"
	"huntsub/huntsub-map-server/api/calendar"
	"huntsub/huntsub-map-server/api/channel"
	"huntsub/huntsub-map-server/api/chatroom"
	"huntsub/huntsub-map-server/api/comment"
	"huntsub/huntsub-map-server/api/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/api/feedback"
	"huntsub/huntsub-map-server/api/library/photo"
	"huntsub/huntsub-map-server/api/library/video"
	"huntsub/huntsub-map-server/api/meeting"
	"huntsub/huntsub-map-server/api/message"
	"huntsub/huntsub-map-server/api/notification"
	"huntsub/huntsub-map-server/api/post"
	"huntsub/huntsub-map-server/api/rank"
	"huntsub/huntsub-map-server/api/user"
	"net/http"
)

type ApiServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewApiServer() *ApiServer {
	var s = &ApiServer{
		ServeMux: http.NewServeMux(),
	}
	s.Handle("/user/", http.StripPrefix("/user", user.NewUserServer()))
	s.Handle("/auth/", http.StripPrefix("/auth", auth.NewAuthServer()))
	s.Handle("/post/", http.StripPrefix("/post", post.NewPostServer()))
	s.Handle("/meeting/", http.StripPrefix("/meeting", meeting.NewMeetingServer()))
	s.Handle("/photo/", http.StripPrefix("/photo", photo.NewPhotoServer()))
	s.Handle("/video/", http.StripPrefix("/video", video.NewVideoServer()))
	s.Handle("/feedback/", http.StripPrefix("/feedback", feedback.NewFeedbackServer()))
	s.Handle("/whfeedback/", http.StripPrefix("/whfeedback", whfeedback.NewWHFeedbackServer()))
	s.Handle("/comment/", http.StripPrefix("/comment", comment.NewCommentServer()))
	s.Handle("/channel/", http.StripPrefix("/channel", channel.NewChannelServer()))
	s.Handle("/chatroom/", http.StripPrefix("/chatroom", chatroom.NewChatRoomServer()))
	s.Handle("/calendar/", http.StripPrefix("/calendar", calendar.NewCalendarServer()))
	s.Handle("/message/", http.StripPrefix("/message", message.NewMessageServer()))
	s.Handle("/like/", http.StripPrefix("/like", like.NewLikeServer()))
	s.Handle("/share/", http.StripPrefix("/share", share.NewShareServer()))
	s.Handle("/follow/", http.StripPrefix("/follow", follow.NewFollowServer()))
	s.Handle("/rank/", http.StripPrefix("/rank", rank.NewRankServer()))
	s.Handle("/report/", http.StripPrefix("/report", report.NewReportServer()))
	s.Handle("/noti/", http.StripPrefix("/noti", notification.NewNotificationManagementServer()))

	return s
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Huntsub-Token"
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", allowedHeaders)
	header.Set("Access-Control-Expose-Headers", "Authorization")
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}

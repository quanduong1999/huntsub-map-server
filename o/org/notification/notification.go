package notification

import "db/mgo"

type NotificationMode map[string]interface{}
type Content map[string]interface{}
type Type string

const (
	LIKE     = Type("like")
	COMMENT  = Type("comment")
	SHARE    = Type("share")
	CALENDAR = Type("calendar")
	RANK     = Type("rank")
	FOLLOW   = Type("follow")
)

type Notification struct {
	UserID   string           `json:"userid"`
	Sender   string           `json:"sender"`
	Avatar   string           `json:"avatar"`
	Type     NotificationMode `json:"type"`
	Ownerid  string           `json:"ownerid"`
	Data     Content          `json:"content"`
	CreateAt int64            `json:"createat"`
	Readed   bool             `json:"readed"`
}

/*******LIKE POST********/
type LikePost struct {
	Type   string `json:"type"`
	ID     string `json:"likeid"`
	PostID string `json:"postid"`
}

/*******LIKE COMMENT********/
type LikeComment struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	PostID string `json:"postid"`
}

/*******LIKE Like Message********/
type LikeMessage struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	PostID string `json:"postid"`
}

/*******Comment POST********/
type CommentPost struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	PostID string `json:"postid"`
}

/*******Repply COMMENT********/
type RepplyComment struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	PostID string `json:"postid"`
}

/*******Share POST********/
type Share struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	PostID string `json:"postid"`
}

/*******Follow********/
type Follow struct {
	Type     string `json:"type"`
	FollowID string `json:"followid"`
}

/*******RANK********/
type Rank struct {
	Type   string `json:"type"`
	RankID string `json:"rankid"`
	Level  string `json:"level"`
}

type NotifiationManagement struct {
	mgo.BaseModel      `bson:",inline"`
	UserID             string `json:"userid" bson:"userid"`
	CalendarNumber     int    `json:"calendar_number" bson:"calendar_number"`
	MessageNumber      int    `json:"message_number" bson:"message_number"`
	NotificationNumber int    `json:"notification_number" bson:"notification_number"`
}

func (s *NotifiationManagement) GetCalendarNumber() int {
	return s.CalendarNumber
}

func (s *NotifiationManagement) GetMessageNumber() int {
	return s.MessageNumber
}

func (s *NotifiationManagement) GetNotificationNumber() int {
	return s.NotificationNumber
}

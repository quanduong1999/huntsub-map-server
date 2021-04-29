package chatroom

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/org/rank"
)

type ChatRoom struct {
	mgo.BaseModel `bson:",inline"`
	Users         []Person    `bson:"users" json:"users"`
	IsGroup       bool        `json:"isgroup" bson:"isgroup"`
	BackGround    string      `json:"background" bson:"background"`
	ColorText     string      `json:"colortext" bson:"colortext"`
	RoomName      string      `json:"roomname" bson:"roomname"`
	LastMessage   LastMessage `json:"lastmessage" bson:"lastmessage"`
}

type Person struct {
	UserID         string        `json:"userid" bson:"userid"`
	Name           string        `json:"name" bson:"name"`
	Avatar         string        `json:"avatar" bson:"avatar"`
	Alias          string        `json:"alias" bson:"alias"`
	Job            string        `json:"job" bson:"job"`
	TimeOut        int64         `json:"timeout" bson:"timeout"`
	RankName       rank.RankName `json:"rankname" bson:"rankname"`
	StarNumber     float64       `json:"starnumber" bson:"starnumber"`
	FeedbackNumber int           `json:"feedbacknumber" bson:"feedbacknumber"`
}

type LastMessage struct {
	Message  string `json:"message" bson:"message"`
	TimeSent int64  `json:"timesent" bson:"timesent"`
}

type ChatRoomForm struct {
	RoomID      string      `json:"roomid"`
	Person      Person      `json:"person"`
	IsGroup     bool        `json:"isgroup"`
	BackGround  string      `json:"background"`
	ColorText   string      `json:"colortext"`
	RoomName    string      `json:"roomname"`
	LastMessage LastMessage `json:"lastmessage"`
}

func (s *ChatRoom) GetUsers() []Person {
	return s.Users
}

func (s *ChatRoom) GetBackGround() string {
	return s.BackGround
}

func (s *ChatRoom) GetIsGroup() bool {
	return s.IsGroup
}

func (s *ChatRoom) GetColorText() string {
	return s.ColorText
}

func (s *ChatRoom) GetRoomName() string {
	return s.RoomName
}

func (s *ChatRoom) GetLastMessage() LastMessage {
	return s.LastMessage
}

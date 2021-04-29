package message

import "db/mgo"

type Message struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string   `json:"userid" bson:"userid"`
	Sender        string   `json:"sender" bson:"sender"`
	Content       Conntent `json:"content" bson:"content"`
	Avatar        string   `json:"avatar" bson:"avatar"`
	RoomID        string   `json:"roomid" bson:"roomid"`
}

type Conntent struct {
	Type     string `json:"type" bson:"type"`
	Conntent string `json:"content" bson:"content"`
	Name     string `json:"name" bson:"name"`
}

func (s *Message) GetSender() string {
	return s.Sender
}

func (s *Message) GetContent() Conntent {
	return s.Content
}

func (s *Message) GetAvatar() string {
	return s.Avatar
}

package session

import (
	"db/mgo"
	"encoding/json"
)

type Session struct {
	mgo.BaseModel `bson:",inline"`
	Username      string `json:"username"`
	UserID        string `json:"user_id"`
	CTime         int64  `json:"ctime"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

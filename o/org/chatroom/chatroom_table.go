package chatroom

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableChatRoom = model.NewTable("huntsub-server", "chatroom", "cr")

func NewPostID() string {
	return TableChatRoom.Next()
}

func (s *ChatRoom) MakeID() string {
	return TableChatRoom.IdMaker.Next()
}

func (b *ChatRoom) Create() error {
	return TableChatRoom.Create(b)
}

func MarkDelete(id string) error {
	return TableChatRoom.MarkDelete(id)
}

func (v *ChatRoom) Update(newValue *ChatRoom) (*ChatRoom, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetUsers(), v.GetUsers()) {
		if len(newValue.GetUsers()) > 0 {
			values["users"] = newValue.GetUsers()
		}
	}
	if newValue.GetBackGround() != v.GetBackGround() && newValue.GetBackGround() != "" {
		values["background"] = newValue.GetBackGround()
	}
	if newValue.GetColorText() != v.GetColorText() && newValue.GetColorText() != "" {
		values["colortext"] = newValue.GetColorText()
	}
	if newValue.GetRoomName() != v.GetRoomName() && newValue.GetRoomName() != "" {
		values["roomname"] = newValue.GetRoomName()
	}
	if !reflect.DeepEqual(newValue.GetLastMessage(), v.GetLastMessage()) {
		values["lastmessage"] = newValue.GetLastMessage()
	}
	err := TableChatRoom.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableChatRoom.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

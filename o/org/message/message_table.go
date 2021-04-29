package message

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableMessage = model.NewTable("huntsub-server", "message", "msg")

func NewPostID() string {
	return TableMessage.Next()
}

func (s *Message) MakeID() string {
	return TableMessage.IdMaker.Next()
}

func (b *Message) Create() error {
	return TableMessage.Create(b)
}

func MarkDelete(id string) error {
	return TableMessage.MarkDelete(id)
}

func (v *Message) Update(newValue *Message) (*Message, error) {
	var values = map[string]interface{}{}

	if newValue.GetSender() != v.GetSender() {
		if newValue.GetSender() != "" {
			values["sender"] = newValue.GetSender()
		}
	}
	if !reflect.DeepEqual(newValue.GetContent(), v.GetContent()) {
		if newValue.Content.Type != "" {
			values["content"] = newValue.GetContent()
		}
	}
	if newValue.GetAvatar() != v.GetAvatar() {
		if newValue.GetAvatar() != "" {
			values["avatar"] = newValue.GetAvatar()
		}
	}

	err := TableMessage.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableMessage.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

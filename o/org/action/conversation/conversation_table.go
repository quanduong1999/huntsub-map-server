package conversation

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableConversation = model.NewTable("huntsub-server", "Conversation", "blo")

func NewConversationID() string {
	return TableConversation.Next()
}

func (s *Conversation) MakeID() string {
	return TableConversation.IdMaker.Next()
}

func (b *Conversation) Create() error {
	return TableConversation.Create(b)
}

func MarkDelete(id string) error {
	return TableConversation.MarkDelete(id)
}

func (v *Conversation) Update(newValue *Conversation) (*Conversation, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetImages(), v.GetImages()) {
		if len(newValue.GetImages()) > 0 {
			values["images"] = newValue.GetImages()
		}
	}

	if !reflect.DeepEqual(newValue.GetFiles(), v.GetFiles()) {
		if len(newValue.GetFiles()) > 0 {
			values["files"] = newValue.GetFiles()
		}
	}

	if !reflect.DeepEqual(newValue.GetLinks(), v.GetLinks()) {
		if len(newValue.GetLinks()) > 0 {
			values["links"] = newValue.GetLinks()
		}
	}

	err := TableConversation.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableConversation.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

package comment

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableComment = model.NewTable("huntsub-server", "comment", "com")

func NewCommentID() string {
	return TableComment.Next()
}

func (s *Comment) MakeID() string {
	return TableComment.IdMaker.Next()
}

func (b *Comment) Create() error {
	return TableComment.Create(b)
}

func MarkDelete(id string) error {
	return TableComment.MarkDelete(id)
}

func (v *Comment) Update(newValue *Comment) (*Comment, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetImages(), v.GetImages()) {
		if newValue.GetImages() != "" {
			values["images"] = newValue.GetImages()
		}
	}

	if !reflect.DeepEqual(newValue.GetText(), v.GetText()) {
		if newValue.GetText() != "" {
			values["text"] = newValue.GetText()
		}
	}
	if !reflect.DeepEqual(newValue.GetName(), v.GetName()) {
		if newValue.GetName() != "" {
			values["name"] = newValue.GetName()
		}
	}
	if !reflect.DeepEqual(newValue.GetAvatar(), v.GetAvatar()) {
		if newValue.GetAvatar() != "" {
			values["avatar"] = newValue.GetAvatar()
		}
	}

	err := TableComment.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableComment.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

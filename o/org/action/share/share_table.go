package share

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableShare = model.NewTable("huntsub-server", "share", "sha")

func NewShareID() string {
	return TableShare.Next()
}

func (s *Share) MakeID() string {
	return TableShare.IdMaker.Next()
}

func (b *Share) Create() error {
	return TableShare.Create(b)
}

func MarkDelete(id string) error {
	return TableShare.MarkDelete(id)
}

func (v *Share) Update(newValue *Share) (*Share, error) {
	var values = map[string]interface{}{}

	if newValue.GetAboutPost() != v.GetAboutPost() {
		if newValue.GetAboutPost() != "" {
			values["aboutpost"] = newValue.GetAboutPost()
		}
	}

	err := TableShare.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableShare.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

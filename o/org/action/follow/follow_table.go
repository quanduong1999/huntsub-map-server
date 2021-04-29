package follow

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableFollow = model.NewTable("huntsub-server", "follow", "fol")

func NewFollowID() string {
	return TableFollow.Next()
}

func (s *Follow) MakeID() string {
	return TableFollow.IdMaker.Next()
}

func (b *Follow) Create() error {
	return TableFollow.Create(b)
}

func MarkDelete(id string) error {
	return TableFollow.MarkDelete(id)
}

func (v *Follow) Update(newValue *Follow) (*Follow, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetIsFollow(), v.GetIsFollow()) {
		values["isfollow"] = newValue.GetIsFollow()
	}

	err := TableFollow.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableFollow.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

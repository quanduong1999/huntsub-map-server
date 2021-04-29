package like

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableLike = model.NewTable("huntsub-server", "like", "lik")

func NewLikeID() string {
	return TableLike.Next()
}

func (s *Like) MakeID() string {
	return TableLike.IdMaker.Next()
}

func (b *Like) Create() error {
	return TableLike.Create(b)
}

func MarkDelete(id string) error {
	return TableLike.MarkDelete(id)
}

func (v *Like) Update(newValue *Like) (*Like, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetAction(), v.GetAction()) {
		if newValue.GetAction() != "" {
			values["action"] = newValue.GetAction()
		}
	}

	err := TableLike.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableLike.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

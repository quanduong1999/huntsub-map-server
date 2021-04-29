package user

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableUserData = model.NewTable("huntsub-server", "userdata", "udt")

func NewFeedBackID() string {
	return TableUserData.Next()
}

func (s *UserData) MakeID() string {
	return TableUserData.IdMaker.Next()
}

func (b *UserData) Create() error {
	return TableUserData.Create(b)
}

func MarkDelete(id string) error {
	return TableUserData.MarkDelete(id)
}

func (v *UserData) Update(newValue *UserData) (*UserData, error) {
	var values = map[string]interface{}{}

	// if newValue.GetPointAVG() != v.GetPointAVG() {
	// 	if newValue.GetPointAVG() != 0 {
	// 		values["pointavg"] = newValue.GetPointAVG()
	// 	}
	// }

	err := TableUserData.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableUserData.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

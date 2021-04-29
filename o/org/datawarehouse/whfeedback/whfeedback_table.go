package whfeedback

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableWHFeedBack = model.NewTable("huntsub-server", "whfeedback", "wfb")

func NewFeedBackID() string {
	return TableWHFeedBack.Next()
}

func (s *WHFeedback) MakeID() string {
	return TableWHFeedBack.IdMaker.Next()
}

func (b *WHFeedback) Create() error {
	return TableWHFeedBack.Create(b)
}

func MarkDelete(id string) error {
	return TableWHFeedBack.MarkDelete(id)
}

func (v *WHFeedback) Update(newValue *WHFeedback) (*WHFeedback, error) {
	var values = map[string]interface{}{}

	if newValue.GetPointAVG() != v.GetPointAVG() {
		if newValue.GetPointAVG() != 0 {
			values["pointavg"] = newValue.GetPointAVG()
		}
	}

	err := TableWHFeedBack.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableWHFeedBack.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

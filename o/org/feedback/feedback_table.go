package feedback

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableFeedBack = model.NewTable("huntsub-server", "feedback", "fb")

func NewFeedBackID() string {
	return TableFeedBack.Next()
}

func (s *Feedback) MakeID() string {
	return TableFeedBack.IdMaker.Next()
}

func (b *Feedback) Create() error {
	return TableFeedBack.Create(b)
}

func MarkDelete(id string) error {
	return TableFeedBack.MarkDelete(id)
}

func (v *Feedback) Update(newValue *Feedback) (*Feedback, error) {
	var values = map[string]interface{}{}

	if newValue.GetPersonName() != v.GetPersonName() {
		if newValue.GetPersonName() != "" {
			values["personname"] = newValue.GetPersonName()
		}
	}

	err := TableFeedBack.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableFeedBack.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

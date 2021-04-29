package report

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableReport = model.NewTable("huntsub-server", "report", "rep")

func NewReportID() string {
	return TableReport.Next()
}

func (s *Report) MakeID() string {
	return TableReport.IdMaker.Next()
}

func (b *Report) Create() error {
	return TableReport.Create(b)
}

func MarkDelete(id string) error {
	return TableReport.MarkDelete(id)
}

func (v *Report) Update(newValue *Report) (*Report, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetLabel(), v.GetLabel()) {
		if newValue.GetLabel() != "" {
			values["label"] = newValue.GetLabel()
		}
	}

	err := TableReport.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableReport.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

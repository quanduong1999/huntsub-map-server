package user

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableUserReport = model.NewTable("huntsub-server", "userreports", "urp")

func NewUserReportID() string {
	return TableUserReport.Next()
}

func (s *UserReport) MakeID() string {
	return TableUserReport.IdMaker.Next()
}

func (b *UserReport) Create() error {
	return TableUserReport.Create(b)
}

func MarkDelete(id string) error {
	return TableUserReport.MarkDelete(id)
}

func (v *UserReport) Update(newValue *UserReport) (*UserReport, error) {
	var values = map[string]interface{}{}

	// if !reflect.DeepEqual(newValue.GetAnalysisFields(), v.GetAnalysisFields()) {
	values["analysisfields"] = newValue.GetAnalysisFields()
	// }

	// if newValue.GetSummary() != v.GetSummary() && newValue.GetSummary() != 0 {
	values["summary"] = newValue.GetSummary()
	// }

	err := TableUserReport.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableUserReport.EnsureIndex(mgo.Index{
	Key:        []string{"name"},
	Background: true,
})

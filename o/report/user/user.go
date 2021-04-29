package user

import (
	"db/mgo"
	"huntsub/huntsub-map-server/x/mlog"
)

var objectUserReportLog = mlog.NewTagLog("object_user_report")

type UserReport struct {
	mgo.BaseModel  `bson:",inline"`
	UserID         string          `json:"userid" bson:"userid"`
	AnalysisFields []AnalysisField `json:"analysisfields" bson:"analysisfields"`
	Summary        float64         `json:"summary" bson:"summary"`
}

type AnalysisField struct {
	Category string  `json:"category" bson:"category"`
	Focus    int     `json:"focus" bson:"focus"`
	UnFocus  int     `json:"unfocus" bson:"unfocus"`
	Percent  float64 `json:"percent" bson:"percent"`
	Sum      int     `json:"sum" bson:"sum"`
}

func (s *UserReport) GetAnalysisFields() []AnalysisField {
	return s.AnalysisFields
}

func (s *UserReport) GetSummary() float64 {
	return s.Summary
}

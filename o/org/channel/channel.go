package channel

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/org/user"
)

type Channel struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string        `json:"userid" bson:"userid"`
	IsActive      bool          `json:"isactive" bson:"isactive"`
	WorkTime      WorkTime      `json:"worktime" bson:"worktime"`
	Job           string        `json:"job" bson:"job"`
	DescribeJob   []string      `json:"describejob" bson:"describejob"`
	Introduction  string        `json:"introduction" bson:"introduction"`
	Location      user.Location `json:"location" bson:"location"`
}

type WorkTime struct {
	Start string `json:"start" bson:"start"`
	End   string `json:"end" bson:"end"`
}

func (s *Channel) GetIsActive() bool {
	return s.IsActive
}

func (s *Channel) GetJob() string {
	return s.Job
}

func (s *Channel) GetDescribeJob() []string {
	return s.DescribeJob
}

func (s *Channel) GetIntroduction() string {
	return s.Introduction
}

func (s *Channel) GetWorkTime() WorkTime {
	return s.WorkTime
}

func (s *Channel) GetLocation() user.Location {
	return s.Location
}

package feedback

import "db/mgo"

type TypeOfFeedBack string

const (
	FeedBackJob    = TypeOfFeedBack("Feedback-Job")
	FeedBackPerson = TypeOfFeedBack("Feedback-Person")
)

type Feedback struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string         `json:"userid" bson:"userid"`
	MeetingID     string         `json:"meetingid" bson:"meetingid"`
	PersonID      string         `json:"personid" bson:"personid"`
	PersonName    string         `json:"personname" bson:"personname"`
	Point         int            `json:"point" bson:"point"`
	Comment       string         `json:"comment" bson:"comment"`
	Type          TypeOfFeedBack `json:"type" bson:"type"`
}

func (s *Feedback) GetPersonName() string {
	return s.PersonName
}

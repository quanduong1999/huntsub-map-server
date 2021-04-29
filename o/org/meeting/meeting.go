package meeting

import (
	"db/mgo"
	"time"
)

type Repeat string
type NotificationTime string

const (
	RepeatDay   = Repeat("Repeat-Day")
	RepeatWeek  = Repeat("Repeat-Week")
	RepeatMonth = Repeat("Repeat-Month")
)

const (
	ThirtyMinus = NotificationTime("30minus")
	SixtyMinus  = NotificationTime("60minus")
)

type Meeting struct {
	mgo.BaseModel    `bson:",inline"`
	WorkerID         string           `json:"workerid" bson:"workerid"`
	BossID           string           `json:"bossid" bson:"bossid"`
	Time             time.Time        `json:"time" bson:"time"`
	Location         string           `json:"location" bson:"location"`
	Repeat           Repeat           `json:"repeat" bson:"repeat"`
	NotificationTime NotificationTime `json:"notificationtime" bson:"notificationtime"`
	Content          string           `json:"content" bson:"content"`
	Title            string           `json:"title" bson:"title"`
	IsFeedback       bool             `json:"isfeedback" bson:"isfeedback"`
}

func (s *Meeting) GetLocation() string {
	return s.Location
}

func (s *Meeting) GetContent() string {
	return s.Content
}

func (s *Meeting) GetTitle() string {
	return s.Title
}

func (s *Meeting) GetIsFeedback() bool {
	return s.IsFeedback
}

func (s *Meeting) GetRepeat() Repeat {
	return s.Repeat
}

func (s *Meeting) GetNotificationTime() NotificationTime {
	return s.NotificationTime
}

package calendar

import "db/mgo"

type Calendar struct {
	mgo.BaseModel    `bson:",inline"`
	UserID           string `json:"userid" bson:"userid"`
	MeetingDay       string `json:"meetingday" bson:"meetingday"`
	MeetingHour      string `json:"meetinghour" bson:"meetinghour"`
	MeetingTime      int    `json:"meeting_time" bson:"meeting_time"`
	ServiceProvider  Person `json:"service_provider" bson:"service_provider"`
	ServiceCaller    Person `json:"service_caller" bson:"service_caller"`
	Title            string `json:"title" bson:"title"`
	Content          string `json:"content" bson:"content"`
	Address          string ` json:"address" bson:"address"`
	RepeatTime       int    `json:"repeat_time" bson:"repeat_time"`
	NotificationTime int    `json:"notification_time" bson:"notification_time"`
	IsConfirm        bool   `json:"is_confirm" bson:"is_confirm"`
	Priority         int    `json:"priority" bson:"priority"`
}

type Person struct {
	Name       string `json:"name" bson:"name"`
	UserID     string `json:"userid" bson:"userid"`
	Avatar     string `json:"avatar" bson:"avatar"`
	FeedbackID string `json:"feedbackid" bson:"feedbackid"`
}

func (s *Calendar) GetMeetingDay() string {
	return s.MeetingDay
}

func (s *Calendar) GetMeetingHour() string {
	return s.MeetingHour
}

func (s *Calendar) GetMeetingTime() int {
	return s.MeetingTime
}

func (s *Calendar) GetServiceProvider() Person {
	return s.ServiceProvider
}

func (s *Calendar) GetServiceCaller() Person {
	return s.ServiceProvider
}

func (s *Calendar) GetTitle() string {
	return s.Title
}

func (s *Calendar) GetContent() string {
	return s.Content
}

func (s *Calendar) GetAddress() string {
	return s.Address
}

func (s *Calendar) GetRepeatTime() int {
	return s.RepeatTime
}

func (s *Calendar) GetNotificationTime() int {
	return s.NotificationTime
}

func (s *Calendar) GetIsConfirm() bool {
	return s.IsConfirm
}

func (s *Calendar) GetPriority() int {
	return s.Priority
}

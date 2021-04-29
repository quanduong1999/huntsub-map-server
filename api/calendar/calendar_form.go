package calendar

import (
	"huntsub/huntsub-map-server/cache/org/feedback"
	"huntsub/huntsub-map-server/o/org/calendar"
)

type CalendarForm struct {
	ID               string          `json:"id"`
	UserID           string          `json:"userid"`
	MeetingDay       string          `json:"meetingday"`
	MeetingHour      string          `json:"meetinghour"`
	MeetingTime      int             `json:"meeting_time"`
	ServiceProvider  calendar.Person `json:"service_provider"`
	ServiceCaller    calendar.Person `json:"service_caller"`
	Title            string          `json:"title"`
	Content          string          `json:"content"`
	Address          string          `json:"address"`
	RepeatTime       int             `json:"repeat_time"`
	NotificationTime int             `json:"notification_time"`
	IsConfirm        bool            `json:"is_confirm"`
	Priority         int             `json:"priority"`
	FeedBackCaller   FeedBack        `json:"feedback_caller"`
	FeedBackProducer FeedBack        `json:"feedback_producer"`
}

type FeedBack struct {
	Content string `json:"content"`
	Star    int    `json:"star"`
}

func NewCalendarForm(c *calendar.Calendar) *CalendarForm {
	var s = &CalendarForm{}

	s.ID = c.ID
	s.UserID = c.UserID
	s.MeetingDay = c.MeetingDay
	s.MeetingHour = c.MeetingHour
	s.MeetingTime = c.MeetingTime
	s.ServiceCaller = c.ServiceCaller
	s.ServiceProvider = c.ServiceProvider
	s.Title = c.Title
	s.Content = c.Content
	s.Address = c.Address
	s.RepeatTime = c.RepeatTime
	s.NotificationTime = c.NotificationTime
	s.IsConfirm = c.IsConfirm
	s.Priority = c.Priority

	fb_caller, err := feedback.Get(c.ServiceCaller.FeedbackID)
	if err != nil {
		/*Client hadnt feefback*/
	} else {
		s.FeedBackCaller.Content = fb_caller.Comment
		s.FeedBackCaller.Star = fb_caller.Point
	}
	fb_producer, err := feedback.Get(c.ServiceProvider.FeedbackID)
	if err != nil {
		/*Client hadnt feefback*/
	} else {
		s.FeedBackProducer.Content = fb_producer.Comment
		s.FeedBackProducer.Star = fb_producer.Point
	}

	return s
}

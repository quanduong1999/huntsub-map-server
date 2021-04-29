package calendar

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableCalendar = model.NewTable("huntsub-server", "calendar", "cld")

func NewPostID() string {
	return TableCalendar.Next()
}

func (s *Calendar) MakeID() string {
	return TableCalendar.IdMaker.Next()
}

func (b *Calendar) Create() error {
	return TableCalendar.Create(b)
}

func MarkDelete(id string) error {
	return TableCalendar.MarkDelete(id)
}

func (v *Calendar) Update(newValue *Calendar) (*Calendar, error) {
	var values = map[string]interface{}{}

	if newValue.GetMeetingDay() != v.GetMeetingDay() && newValue.GetMeetingDay() != "" {
		values["mettingday"] = newValue.GetMeetingDay()
	}
	if newValue.GetMeetingHour() != v.GetMeetingHour() && newValue.GetMeetingHour() != "" {
		values["mettinghour"] = newValue.GetMeetingHour()
	}
	if !reflect.DeepEqual(newValue.GetServiceProvider(), v.GetServiceProvider()) {
		values["service_provider"] = newValue.GetServiceProvider()
	}
	if !reflect.DeepEqual(newValue.GetServiceCaller(), v.GetServiceCaller()) {
		values["service_caller"] = newValue.GetServiceCaller()
	}
	if newValue.GetMeetingTime() != v.GetMeetingTime() {
		values["meeting_time"] = newValue.GetMeetingTime()
	}
	if newValue.GetTitle() != v.GetTitle() && newValue.Title != "" {
		values["title"] = newValue.GetTitle()
	}
	if newValue.GetContent() != v.GetContent() && newValue.Content != "" {
		values["content"] = newValue.GetContent()
	}

	if !reflect.DeepEqual(newValue.GetAddress(), v.GetAddress()) && newValue.GetAddress() != "" {
		values["address"] = newValue.GetAddress()
	}

	if newValue.GetRepeatTime() != v.GetRepeatTime() {
		values["repeat_time"] = newValue.GetRepeatTime()
	}

	if newValue.GetNotificationTime() != v.GetNotificationTime() {
		values["notification_time"] = newValue.GetNotificationTime()
	}
	if newValue.GetIsConfirm() != v.GetIsConfirm() {
		values["is_confirm"] = newValue.GetIsConfirm()
	}
	if newValue.GetPriority() != 0 {
		values["priority"] = newValue.GetPriority()
	}

	err := TableCalendar.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableCalendar.EnsureIndex(mgo.Index{
	Key:        []string{"meeting_time"},
	Background: true,
})

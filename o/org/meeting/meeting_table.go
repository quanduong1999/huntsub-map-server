package meeting

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableMeeting = model.NewTable("huntsub-server", "meeting", "fb")

func NewMeetingID() string {
	return TableMeeting.Next()
}

func (s *Meeting) MakeID() string {
	return TableMeeting.IdMaker.Next()
}

func (b *Meeting) Create() error {
	return TableMeeting.Create(b)
}

func MarkDelete(id string) error {
	return TableMeeting.MarkDelete(id)
}

func (v *Meeting) Update(newValue *Meeting) (*Meeting, error) {
	var values = map[string]interface{}{}

	if newValue.GetLocation() != v.GetLocation() {
		if newValue.GetLocation() != "" {
			values["location"] = newValue.GetLocation()
		}
	}
	if newValue.GetContent() != v.GetContent() {
		if newValue.GetContent() != "" {
			values["content"] = newValue.GetContent()
		}
	}
	if newValue.GetTitle() != v.GetTitle() {
		if newValue.GetTitle() != "" {
			values["title"] = newValue.GetTitle()
		}
	}
	if newValue.GetIsFeedback() != v.GetIsFeedback() {
		values["isfeedback"] = newValue.GetIsFeedback()
	}
	if newValue.GetRepeat() != v.GetRepeat() {
		if newValue.GetRepeat() != "" {
			values["repeat"] = newValue.GetRepeat()
		}
	}
	if newValue.GetNotificationTime() != v.GetNotificationTime() {
		values["notificationtime"] = newValue.GetNotificationTime()
	}

	err := TableMeeting.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableMeeting.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

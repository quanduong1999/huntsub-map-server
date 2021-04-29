package notification

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
)

var TableNotifiationManagement = model.NewTable("huntsub-server", "notification", "not")

func NewPostID() string {
	return TableNotifiationManagement.Next()
}

func (s *NotifiationManagement) MakeID() string {
	return TableNotifiationManagement.IdMaker.Next()
}

func (b *NotifiationManagement) Create() error {
	var _ = TableNotifiationManagement.EnsureAddressIndex()
	return TableNotifiationManagement.Create(b)
}

func MarkDelete(id string) error {
	return TableNotifiationManagement.MarkDelete(id)
}

func (v *NotifiationManagement) Update(newValue *NotifiationManagement) (*NotifiationManagement, error) {
	var values = map[string]interface{}{}

	if newValue.GetCalendarNumber() != v.GetCalendarNumber() {
		values["calendar_number"] = newValue.GetCalendarNumber()
	}

	if newValue.GetMessageNumber() != v.GetMessageNumber() {
		values["message_number"] = newValue.GetMessageNumber()
	}

	if newValue.GetNotificationNumber() != v.GetNotificationNumber() {
		values["notification_number"] = newValue.GetNotificationNumber()
	}

	err := TableNotifiationManagement.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableNotifiationManagement.EnsureIndex(mgo.Index{
	Key:        []string{"userid"},
	Background: true,
})

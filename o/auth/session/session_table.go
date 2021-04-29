package session

import (
	"huntsub/huntsub-map-server/o/model"
)

var TableSession = model.NewTable("huntsub-server", "session", "ses")

func (b *Session) Create() error {
	return TableSession.Create(b)
}

func MarkDelete(id string) error {
	return TableSession.MarkDelete(id)
}

func (v *Session) Update(newValue *Session) error {
	return TableSession.UnsafeUpdateByID(v.ID, newValue)
}

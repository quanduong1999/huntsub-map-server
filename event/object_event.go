package event

import (
	"huntsub/huntsub-map-server/event/shared/oev"
	"mrw/event"
)

type objectEventSource struct {
	ev *event.Hub
}

func (s *objectEventSource) emit(action oev.ObjectActionName, data interface{}) {
	var o = oev.NewObjectEvent(action, data)
	s.ev.Emit(o)
}

func (s *objectEventSource) EmitCreate(data interface{}) {
	s.emit(oev.ObjectActionCreate, data)
}

func (s *objectEventSource) EmitUpdate(data interface{}) {
	s.emit(oev.ObjectActionUpdate, data)
}

func (s *objectEventSource) EmitMarkDelete(data interface{}) {
	s.emit(oev.ObjectActionMarkDelete, data)
}

func (s *objectEventSource) OnEvent() (event.Line, event.Cancel) {
	return s.ev.NewLine()
}

func (s *objectEventSource) EmitConfirm(data interface{}) {
	s.emit(oev.ObjectActionConfirm, data)
}

var ObjectEventSource = &objectEventSource{
	ev: event.NewHub(event.LargeHub),
}

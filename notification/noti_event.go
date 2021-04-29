package notification

import (
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/notification/action/follow"
	"huntsub/huntsub-map-server/notification/action/like"
	"huntsub/huntsub-map-server/notification/action/share"
	"huntsub/huntsub-map-server/notification/calendar"
	"huntsub/huntsub-map-server/notification/feedback"
	"huntsub/huntsub-map-server/notification/meeting"
	"huntsub/huntsub-map-server/notification/post"
	"mrw/event"
)

var notiPush = event.NewHub(event.LargeHub)
var notiReady = event.NewHub(event.SmallHub)

func OnChange() (event.Line, event.Cancel) {
	return notiPush.NewLine()
}

func ready() {
	notiReady.Emit(true)
	notificationLog.Infof(0, "notification pusher is ready")
}

func Wait() {
	v, _ := notiReady.Value().(bool)
	if v {
		return
	}
	ready, cancel := notiReady.NewLine()
	defer cancel()
	<-ready
}

func handleEvent(v interface{}) {
	obj, ok := v.(*oev.ObjectEvent)
	if !ok {
		return
	}
	notificationLog.Infof(0, "Handle Notification Pusher  %s", v)
	switch obj.Category {
	case oev.ObjectCategoryPost:
		post.Handle(obj)
	case oev.ObjectCategoryMeeting:
		meeting.Handle(obj)
	case oev.ObjectCategoryActionLike:
		like.Handle(obj)
	case oev.ObjectCategoryActionShare:
		share.Handle(obj)
	case oev.ObjectCategoryActionFollow:
		follow.Handle(obj)
	case oev.ObjectCategoryCalendar:
		calendar.Handle(obj)
	case oev.ObjectCategoryFeedback:
		feedback.Handle(obj)
	default:
		return
	}
	//After handle event --> emit
	notiPush.Emit(v)
}

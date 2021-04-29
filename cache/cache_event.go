package cache

import (
	"huntsub/huntsub-map-server/cache/org/action/follow"
	"huntsub/huntsub-map-server/cache/org/action/like"
	"huntsub/huntsub-map-server/cache/org/calendar"
	"huntsub/huntsub-map-server/cache/org/chatroom"
	"huntsub/huntsub-map-server/cache/org/comment"
	"huntsub/huntsub-map-server/cache/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/cache/org/feedback"
	"huntsub/huntsub-map-server/cache/org/library/photo"
	"huntsub/huntsub-map-server/cache/org/library/video"
	"huntsub/huntsub-map-server/cache/org/meeting"
	"huntsub/huntsub-map-server/cache/org/notification"
	"huntsub/huntsub-map-server/cache/org/post"
	userReport "huntsub/huntsub-map-server/cache/org/report/user"
	"huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/event/shared/oev"
	"mrw/event"
)

var cacheChange = event.NewHub(event.LargeHub)
var cacheReady = event.NewHub(event.SmallHub)

func OnChange() (event.Line, event.Cancel) {
	return cacheChange.NewLine()
}

func ready() {
	cacheReady.Emit(true)
	cacheLog.Infof(0, "cache is ready")
}

func Wait() {
	v, _ := cacheReady.Value().(bool)
	if v {
		return
	}
	ready, cancel := cacheReady.NewLine()
	defer cancel()
	<-ready
}

func handleEvent(v interface{}) {
	obj, ok := v.(*oev.ObjectEvent)
	if !ok {
		return
	}
	cacheLog.Infof(0, "Handle Event %s", v)
	switch obj.Category {
	case oev.ObjectCategoryUser:
		user.Handle(obj)
	case oev.ObjectCategoryPost:
		post.Handle(obj)
	case oev.ObjectCategoryMeeting:
		meeting.Handle(obj)
	case oev.ObjectCategoryPhoto:
		photo.Handle(obj)
	case oev.ObjectCategoryVideo:
		video.Handle(obj)
	case oev.ObjectCategoryFeedback:
		feedback.Handle(obj)
	case oev.ObjectCategoryWHFeedback:
		whfeedback.Handle(obj)
	case oev.ObjectCategoryComment:
		comment.Handle(obj)
	case oev.ObjectCategoryChatroom:
		chatroom.Handle(obj)
	case oev.ObjectCategoryCalendar:
		calendar.Handle(obj)
	case oev.ObjectCategoryActionLike:
		like.Handle(obj)
	case oev.ObjectCategoryActionComment:
		comment.Handle(obj)
	case oev.ObjectCategoryActionFollow:
		follow.Handle(obj)
	case oev.ObjectCategoryNotification:
		notification.Handle(obj)
	case oev.ObjectCategoryUserReport:
		userReport.Handle(obj)
	}
	//After handle event --> emit
	cacheChange.Emit(v)
}

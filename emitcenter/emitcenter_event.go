package emitcenter

import (
	"huntsub/huntsub-map-server/emitcenter/action/comment"
	"huntsub/huntsub-map-server/emitcenter/action/like"
	"huntsub/huntsub-map-server/emitcenter/action/share"
	"huntsub/huntsub-map-server/emitcenter/feedback"
	"huntsub/huntsub-map-server/emitcenter/meeting"
	"huntsub/huntsub-map-server/emitcenter/post"
	"huntsub/huntsub-map-server/emitcenter/user"
	"huntsub/huntsub-map-server/event/shared/oev"
	"mrw/event"
)

var emitCenter = event.NewHub(event.LargeHub)
var emitcenterReady = event.NewHub(event.SmallHub)

func OnChange() (event.Line, event.Cancel) {
	return emitCenter.NewLine()
}

func ready() {
	emitcenterReady.Emit(true)
	emitcenterLog.Infof(0, "emitcenter handles signal is ready")
}

func Wait() {
	v, _ := emitcenterReady.Value().(bool)
	if v {
		return
	}
	ready, cancel := emitcenterReady.NewLine()
	defer cancel()
	<-ready
}

func handleEvent(v interface{}) {
	obj, ok := v.(*oev.ObjectEvent)
	if !ok {
		return
	}
	emitcenterLog.Infof(0, "Handle EmitCenter  %s", v)
	switch obj.Category {
	case oev.ObjectCategoryPost:
		post.Handle(obj)
	case oev.ObjectCategoryMeeting:
		meeting.Handle(obj)
	case oev.ObjectCategoryActionLike:
		like.Handle(obj)
	case oev.ObjectCategoryActionShare:
		share.Handle(obj)
	case oev.ObjectCategoryUser:
		user.Handle(obj)
	case oev.ObjectCategoryFeedback:
		feedback.Handle(obj)
	case oev.ObjectCategoryComment:
		comment.Handle(obj)
	default:
		return
	}
	//After handle event --> emit
	emitCenter.Emit(v)
}

package feedback

import (
	"http/web"
	cache_calendar "huntsub/huntsub-map-server/cache/org/calendar"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/feedback"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var obj = event.ObjectEventSource
var notiPostLog = mlog.NewTagLog("noti_feedback")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	fb, _ := o.Data.(*feedback.Feedback)
	notiPostLog.Infof(0, "Handle Notification Feedback %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return HandleFeedbackCreating(fb)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func HandleFeedbackCreating(fb *feedback.Feedback) error {
	/* Increase point of Rank User*/
	ranUsr, err := rank.GetRank(map[string]interface{}{
		"dtime":  0,
		"userid": fb.UserID,
	})
	newRan := ranUsr
	newRan.Radius++
	res, err := ranUsr.Update(newRan)
	web.AssertNil(err)
	obj.EmitUpdate(res)
	/* Handle point of Rank to Person was be feedback*/
	ranPer, err := rank.GetRank(map[string]interface{}{
		"dtime":  0,
		"userid": fb.PersonID,
	})

	newRanPer := ranPer
	switch fb.Point {
	case 1:
		newRanPer.Radius -= 2
		break
	case 2:
		newRanPer.Radius -= 1
		break
	case 3:
		break
	case 4:
		newRanPer.Radius += 1
		break
	case 5:
		newRanPer.Radius += 2
		break
	default:
		break
	}
	res, err = ranPer.Update(newRanPer)
	web.AssertNil(err)
	obj.EmitUpdate(res)

	/*Handle status of Calendar*/
	cal, err := cache_calendar.Get(fb.MeetingID)
	cal.Priority++
	web.AssertNil(err)
	resCal, err := cal.Update(cal)
	obj.EmitUpdate(resCal)
	return err
}

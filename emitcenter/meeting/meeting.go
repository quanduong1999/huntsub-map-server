package meeting

import (
	"fmt"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/meeting"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var notiMeetingtLog = mlog.NewTagLog("noti_post")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	meet, _ := o.Data.(*meeting.Meeting)
	fmt.Println(meet)
	notiMeetingtLog.Infof(0, "Handle Notification Meeting %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return nil
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}
	return nil
}

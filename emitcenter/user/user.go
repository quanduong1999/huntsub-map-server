package user

import (
	"http/web"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/o/org/user"
	userReport "huntsub/huntsub-map-server/o/report/user"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var notiUserLog = mlog.NewTagLog("noti_user")
var obj = event.ObjectEventSource

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	u, _ := o.Data.(*user.User)
	notiUserLog.Infof(0, "Handle Notification User %s", o.Action)
	switch o.Action {
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionUpdate:
		return HandleActionUpdate(u)
	case oev.ObjectActionCreate:
		return HandleActionCreate(u)
	}

	return nil
}

func HandleActionCreate(u *user.User) error {
	/*Create User's channel */
	var cn = &channel.Channel{}
	cn.UserID = u.GetID()
	cn.Location = u.Location
	cn.Job = u.Job
	web.AssertNil(cn.Create())
	obj.EmitCreate(cn)

	/*Create User's Notification management */
	var noti = &notification.NotifiationManagement{}
	noti.UserID = u.GetID()
	web.AssertNil(noti.Create())
	obj.EmitCreate(noti)

	/*Create User's Warehouse Feedback*/
	var whfb = &whfeedback.WHFeedback{}
	whfb.UserID = u.GetID()
	web.AssertNil(whfb.Create())
	obj.EmitCreate(whfb)

	/*Create User's rank*/
	var ran = &rank.Rank{}
	ran.UserID = u.GetID()
	ran.RankName = rank.NguoiMoi_I
	ran.Location = u.Location
	ran.Job = u.Job
	web.AssertNil(ran.Create())
	obj.EmitCreate(ran)

	/*Create User's report*/
	var urp = &userReport.UserReport{}
	urp.UserID = u.GetID()
	web.AssertNil(urp.Create())
	obj.EmitCreate(urp)
	return nil
}

func HandleActionUpdate(u *user.User) error {
	/*Update current location of channel when user updating*/
	var cn = &channel.Channel{}
	cn.Location = u.Location
	c, err := channel.GetChannel(map[string]interface{}{
		"dtime":  0,
		"userid": u.GetID(),
	})
	cn.IsActive = c.IsActive
	web.AssertNil(err)
	res, err := c.Update(cn)
	web.AssertNil(err)
	obj.EmitUpdate(res)
	return nil
}

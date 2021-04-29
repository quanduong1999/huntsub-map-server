package calendar

import (
	"db/firebase"
	"fmt"
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/config/cons"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/calendar"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"time"
	"util/runtime"
)

var notiCalendarLog = mlog.NewTagLog("noti_calendar")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	cl, _ := o.Data.(*calendar.Calendar)
	notiCalendarLog.Infof(0, "Handle Notification Calendar %s", o.Action)
	switch o.Action {
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionConfirm:
		return HandleActionConfirm(cl)
	case oev.ObjectActionCreate:
		return HandleCreateCalender(cl)
	}

	return nil
}

func HandleCreateCalender(ca *calendar.Calendar) error {
	var noti = &notification.Notification{}
	noti.UserID = ca.UserID
	noti.CreateAt = time.Now().Local().Unix()
	if ca.UserID == ca.ServiceProvider.UserID {
		noti.Ownerid = ca.ServiceCaller.UserID
	}
	var data = notification.NotificationMode{}
	data["type"] = "calendar"
	data["calendarid"] = ca.ID

	noti.Type = data
	u, err := cache_user.Get(noti.UserID)
	web.AssertNil(err)
	owner, err := cache_user.Get(noti.Ownerid)
	web.AssertNil(err)
	noti.Sender = u.Name
	noti.Avatar = u.Avatar
	client, _ := firebase.NewClientFirebase()
	_, _, err = client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding calendar: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user*/
	if owner.StatusActive.Online {

	} else {
		message := u.Name + "DA TAO LICH HEN VOI BAN"
		common.SendNotification(owner.ExpoToken, message)
	}
	return err
}

func HandleActionConfirm(ca *calendar.Calendar) error {
	var noti = &notification.Notification{}
	noti.UserID = ca.UserID
	noti.CreateAt = time.Now().Local().Unix()
	if ca.UserID == ca.ServiceProvider.UserID {
		noti.Ownerid = ca.ServiceCaller.UserID
	}
	var data = notification.NotificationMode{}
	data["type"] = cons.ConfirmCalendar
	data["calendarid"] = ca.ID

	noti.Type = data
	u, err := cache_user.Get(noti.UserID)
	web.AssertNil(err)
	owner, err := cache_user.Get(noti.Ownerid)
	web.AssertNil(err)
	noti.Sender = u.Name
	noti.Avatar = u.Avatar
	client, _ := firebase.NewClientFirebase()
	_, _, err = client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding confirm calendar: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user*/
	if owner.StatusActive.Online {

	} else {
		message := fmt.Sprintf("%s muốn huỷ lịch hẹn với bạn", u.Name)
		common.SendNotification(owner.ExpoToken, message)
	}
	return err
}

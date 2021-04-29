package feedback

import (
	"db/firebase"
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/feedback"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"time"
	"util/runtime"
)

var notiPostLog = mlog.NewTagLog("noti_feedback")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	fb, _ := o.Data.(*feedback.Feedback)
	notiPostLog.Infof(0, "Handle Notification Feedback %s", o.Action)
	switch o.Action {
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionReport:
		return nil
	case oev.ObjectActionCreate:
		return Add(fb)
	}

	return nil
}

func Add(fb *feedback.Feedback) error {
	var noti = &notification.Notification{}
	noti.UserID = fb.UserID
	noti.Ownerid = fb.UserID
	noti.CreateAt = time.Now().Local().Unix()
	var data = notification.NotificationMode{}
	data["type"] = "feedback"
	data["feedbackid"] = fb.ID

	noti.Type = data

	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding like: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user*/
	user, err := cache_user.Get(fb.UserID)
	web.AssertNil(err)
	owner, err := cache_user.Get(fb.PersonID)
	web.AssertNil(err)
	if owner.StatusActive.Online {

	} else {
		message := user.Name + "Đã gửi một feedback cho bạn"
		common.SendNotification(owner.ExpoToken, message)
	}
	return err
}

package rank

import (
	"db/firebase"
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"time"
	"util/runtime"
)

var notiPostLog = mlog.NewTagLog("noti_post")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	ran, _ := o.Data.(*rank.Rank)
	notiPostLog.Infof(0, "Handle Notification Rank %s", o.Action)
	switch o.Action {
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionReport:
		return nil
	case oev.ObjectActionCreate:
		return Add(ran)
	}

	return nil
}

func Add(ran *rank.Rank) error {
	var noti = &notification.Notification{}
	noti.UserID = ran.UserID
	noti.Ownerid = ran.UserID
	noti.CreateAt = time.Now().Local().Unix()
	var data = notification.NotificationMode{}
	data["type"] = "rank"
	data["rankid"] = ran.ID

	noti.Type = data

	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding like: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user*/
	owner, err := cache_user.Get(ran.UserID)
	web.AssertNil(err)
	if owner.StatusActive.Online {

	} else {
		message := "CHUC MUNG BAN DA THANG HANG" + owner.Name
		common.SendNotification(owner.ExpoToken, message)
	}
	return err
}

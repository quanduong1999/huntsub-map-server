package follow

import (
	"db/firebase"
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/follow"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"util/runtime"
)

var notiLikeLog = mlog.NewTagLog("noti_like")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	fol, _ := o.Data.(*follow.Follow)
	notiLikeLog.Infof(0, "Handle Notification Follow %s", o.Action)

	switch o.Action {
	case oev.ObjectActionCreate:
		return Add(fol)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func Add(fol *follow.Follow) error {
	var noti = &notification.Notification{}
	noti.UserID = fol.UserID
	noti.Ownerid = fol.PersonID
	var data = notification.NotificationMode{}
	data["type"] = "follow"
	data["followid"] = fol.ID

	noti.Type = data

	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding follow: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user*/
	owner, err := cache_user.Get(fol.PersonID)
	web.AssertNil(err)
	userFollow, err := cache_user.Get(fol.UserID)
	web.AssertNil(err)
	if owner.StatusActive.Online {

	} else {
		message := userFollow.Name + " Đã follow " + owner.Name
		common.SendNotification(owner.ExpoToken, message)
	}
	return err
}

package like

import (
	"db/firebase"
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/like"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"time"
	"util/runtime"
)

var notiLikeLog = mlog.NewTagLog("noti_like")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	lik, _ := o.Data.(*like.Like)
	notiLikeLog.Infof(0, "Handle Notification Like %s", o.Action)

	switch o.Action {
	case oev.ObjectActionCreate:
		return Add(lik)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func Add(lik *like.Like) error {
	if lik.UserID == lik.OwnerID {
		return nil
	}
	var noti = &notification.Notification{}
	noti.UserID = lik.UserID
	noti.Ownerid = lik.OwnerID
	noti.CreateAt = time.Now().Local().Unix()
	var data = notification.NotificationMode{}
	data["type"] = "like"
	data["likeid"] = lik.ID
	data["postid"] = lik.PostID
	noti.Type = data

	owner, err := cache_user.Get(lik.OwnerID)
	web.AssertNil(err)
	userLike, err := cache_user.Get(lik.UserID)
	web.AssertNil(err)
	noti.Sender = userLike.Name
	noti.Avatar = userLike.Avatar

	client, _ := firebase.NewClientFirebase()
	_, _, err = client.ClientStore.Collection("notifications").Add(client.Ctx, noti)
	if err != nil {
		log.Fatalf("Failed adding like: %v", err)
		return err
	}
	defer client.ClientStore.Close()

	/*Check Status Active user for send Notication*/
	if owner.StatusActive.Online && lik.OwnerID == lik.UserID {

	} else {
		message := userLike.Name + "Đã like bài viết của bạn" + owner.Name
		common.SendNotification(owner.ExpoToken, message)
	}

	return err
}

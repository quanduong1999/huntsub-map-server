package post

import (
	"db/firebase"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"util/runtime"
)

var notiPostLog = mlog.NewTagLog("noti_post")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	pos, _ := o.Data.(*post.Post)
	notiPostLog.Infof(0, "Handle Notification Post %s", o.Action)
	switch o.Action {
	case oev.ObjectActionLike:
		return hadnleLike(pos)
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionReport:
		return nil
	case oev.ObjectActionCreate:
		return Add(pos)
	}

	return nil
}

func hadnleLike(pos *post.Post) error {
	// Todo Like action
	return nil
}

func Add(pos *post.Post) error {
	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("posts").Add(client.Ctx, pos)
	if err != nil {
		log.Fatalf("Failed adding post: %v", err)
		return err
	}
	defer client.ClientStore.Close()
	return err
}

package share

import (
	"db/firebase"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/share"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"util/runtime"
)

var notiShareLog = mlog.NewTagLog("noti_like")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	lik, _ := o.Data.(*share.Share)
	notiShareLog.Infof(0, "Handle Notification Like %s", o.Action)

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

func Add(sha *share.Share) error {
	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("shares").Add(client.Ctx, sha)
	if err != nil {
		log.Fatalf("Failed adding like: %v", err)
		return err
	}
	defer client.ClientStore.Close()
	return err
}

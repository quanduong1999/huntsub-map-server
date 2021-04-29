package message

import (
	"db/firebase"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/message"
	"huntsub/huntsub-map-server/x/mlog"
	"log"
	"util/runtime"
)

var notiMessageLog = mlog.NewTagLog("noti_message")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	msg, _ := o.Data.(*message.Message)
	notiMessageLog.Infof(0, "Handle Notification Message %s", o.Action)
	switch o.Action {
	case oev.ObjectActionLike:
		return hadnleLike(msg)
	case oev.ObjectActionShare:
		return nil
	case oev.ObjectActionReport:
		return nil
	case oev.ObjectActionCreate:
		return Add(msg)
	}

	return nil
}

func hadnleLike(pos *message.Message) error {
	// Todo Like action
	return nil
}

func Add(pos *message.Message) error {
	client, _ := firebase.NewClientFirebase()
	_, _, err := client.ClientStore.Collection("message").Add(client.Ctx, pos)
	if err != nil {
		log.Fatalf("Failed adding message: %v", err)
		return err
	}
	defer client.ClientStore.Close()
	return err
}

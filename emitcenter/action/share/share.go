package share

import (
	"http/web"
	cache_post "huntsub/huntsub-map-server/cache/org/post"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/share"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var emitcenterShareLog = mlog.NewTagLog("emitcenter_share")
var obj = event.ObjectEventSource

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	sha, _ := o.Data.(*share.Share)
	emitcenterShareLog.Infof(0, "Handle Notification Like %s", o.Action)

	switch o.Action {
	case oev.ObjectActionCreate:
		return HandleShareCreating(sha)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func HandleShareCreating(sha *share.Share) error {
	//Update Post.Share++
	p, err := cache_post.Get(sha.PostID)
	web.AssertNil(err)
	var newPost = &post.Post{}
	newPost.Share = p.Share + 1
	res, err := p.Update(newPost)
	web.AssertNil(err)
	obj.EmitUpdate(res)

	/*Update Share total of Rank*/
	ran, err := rank.GetRank(map[string]interface{}{
		"dtime":  0,
		"userid": sha.OwnerID,
	})
	newRan := ran
	newRan.ShareNumber++
	resRan, err := ran.Update(newRan)
	web.AssertNil(err)
	obj.EmitUpdate(resRan)
	return err
}

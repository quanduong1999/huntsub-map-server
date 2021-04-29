package comment

import (
	"http/web"
	cache_post "huntsub/huntsub-map-server/cache/org/post"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/comment"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var emitcenterShareLog = mlog.NewTagLog("emitcenter_share")
var obj = event.ObjectEventSource

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	com, _ := o.Data.(*comment.Comment)
	emitcenterShareLog.Infof(0, "Handle Notification Like %s", o.Action)

	switch o.Action {
	case oev.ObjectActionCreate:
		return HandleCommentCreating(com)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func HandleCommentCreating(com *comment.Comment) error {
	//Update Post.Comment++
	p, err := cache_post.Get(com.PostID)
	web.AssertNil(err)
	var newPost = &post.Post{}
	newPost.Comment = p.Comment + 1
	res, err := p.Update(newPost)
	web.AssertNil(err)
	obj.EmitUpdate(res)

	/*Update Comment total of Rank*/

	return err
}

package like

import (
	"http/web"
	cache_post "huntsub/huntsub-map-server/cache/org/post"
	"huntsub/huntsub-map-server/common"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/like"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/o/report/user"
	"huntsub/huntsub-map-server/x/mlog"
	"util/runtime"
)

var emitcenterLikeLog = mlog.NewTagLog("emitcenter_like")
var obj = event.ObjectEventSource

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	lik, _ := o.Data.(*like.Like)
	emitcenterLikeLog.Infof(0, "Handle Emit Center Like %s", o.Action)

	switch o.Action {
	case oev.ObjectActionCreate:
		return HandleLikeCreating(lik)
	case oev.ObjectActionUpdate:
		return nil
	case oev.ObjectActionMarkDelete:
		return nil
	}

	return nil
}

func HandleLikeCreating(lik *like.Like) error {
	//Update Post.Like++
	p, err := cache_post.Get(lik.PostID)
	web.AssertNil(err)
	var newPost = &post.Post{}
	newPost.Like = p.Like + 1
	res, err := p.Update(newPost)
	web.AssertNil(err)
	obj.EmitUpdate(res)

	/*Update Like total of Rank*/
	ran, err := rank.GetRank(map[string]interface{}{
		"dtime":  0,
		"userid": lik.OwnerID,
	})
	newRan := ran
	newRan.LikeNumber++
	resRan, err := ran.Update(newRan)
	web.AssertNil(err)
	obj.EmitUpdate(resRan)
	/*Handler to analysis fields client focus*/
	urp, err := user.GetUserReport(map[string]interface{}{
		"dtime":  0,
		"userid": lik.UserID,
	})
	web.AssertNil(err)
	newUrp := urp
	index := common.CheckExist(p.Category, newUrp)
	if index == -1 {
		//add new Analysis
		var analy = user.AnalysisField{}
		analy.Category = p.Category
		if lik.Action == like.ActionDisLike {
			analy.UnFocus = 1
		} else {
			analy.Focus = 1
		}
		analy.Sum = 1
		newUrp.AnalysisFields = append(newUrp.AnalysisFields, analy)
	} else {
		if lik.Action == like.ActionDisLike {
			newUrp.AnalysisFields[index].UnFocus += 1
			newUrp.AnalysisFields[index].Percent = float64(newUrp.AnalysisFields[index].Focus-newUrp.AnalysisFields[index].UnFocus) / (newUrp.Summary + 1)
			newUrp.AnalysisFields[index].Sum += 1
		} else {
			newUrp.AnalysisFields[index].Focus += 1
			newUrp.AnalysisFields[index].Percent = float64(newUrp.AnalysisFields[index].Focus-newUrp.AnalysisFields[index].UnFocus) / (newUrp.Summary + 1)
			newUrp.AnalysisFields[index].Sum += 1
		}

	}
	newUrp.Summary++
	_urp, err := urp.Update(newUrp)
	web.AssertNil(err)
	obj.EmitUpdate(_urp)
	return err
}

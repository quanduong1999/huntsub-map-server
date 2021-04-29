package comment

import (
	"http/web"
	"huntsub/huntsub-map-server/o/org/comment"

	"gopkg.in/mgo.v2/bson"
)

type CommentForm struct {
	CommentRoot  *comment.Comment   `json:"commentroot"`
	CommentChild []*comment.Comment `json:"commentchild"`
}

func NewCommentForm(postID, _type string, skip, limit int) ([]*CommentForm, error) {
	var CF = []*CommentForm{}

	//get limit Comment Root record
	var queryRoot = map[string]interface{}{
		"dtime":         0,
		"postid":        postID,
		"commentrootid": bson.M{"$eq": ""},
	}
	res, ok := comment.GetByPaginationAd(queryRoot, skip, limit, _type)
	web.AssertNil(ok)

	for _, u := range res {
		var s = &CommentForm{}
		s.CommentRoot = u
		queryChild := map[string]interface{}{
			"dtime":         0,
			"postid":        postID,
			"commentrootid": u.GetID(),
		}
		childs, err := comment.GetByPaginationAd(queryChild, skip, limit, _type)
		web.AssertNil(err)
		for _, c := range childs {
			s.CommentChild = append(s.CommentChild, c)
		}

		CF = append(CF, s)
	}

	return CF, nil
}

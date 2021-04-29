package like

import "db/mgo"

type Action string

const (
	ActionLike    = Action("like")
	ActionDisLike = Action("dislike")
	ActionHaha    = Action("haha")
	ActionAngry   = Action("angry")
	ActionCry     = Action("cry")
	ActionSad     = Action("sad")
	ActionSpice   = Action("spice") //cay
)

type Like struct {
	mgo.BaseModel `bson:",inline"`
	PostID        string `json:"postid" bson:"postid"`
	CommentID     string `json:"commentid" bson:"commentid"`
	Action        Action `json:"action" bson:"action"`
	UserID        string `json:"userid" bson:"userid"`
	OwnerID       string `json:"ownerid" bson:"ownerid"`
}

func (s *Like) GetAction() Action {
	return s.Action
}

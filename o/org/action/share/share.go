package share

import "db/mgo"

type Share struct {
	mgo.BaseModel `bson:",inline"`
	PostID        string `json:"postid" bson:"postid"`
	AboutPost     string `json:"aboutpost" bson:"aboutpost"`
	UserID        string `json:"userid" bson:"userid"`
	OwnerID       string `json:"ownerid" bson:"ownerid"`
}

func (s *Share) GetAboutPost() string {
	return s.AboutPost
}

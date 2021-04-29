package comment

import "db/mgo"

type Comment struct {
	mgo.BaseModel `bson:",inline"`
	PostID        string `json:"postid" bson:"postid"`
	Text          string `json:"text" bson:"text"`
	Image         string `json:"image" bson:"image"`
	CommentRootID string `json:"commentrootid" bson:"commentrootid"`
	UserID        string `json:"userid" bson:"userid"`
	Name          string `json:"name" bson:"name"`
	Avatar        string `json:"avatar" bson:"avatar"`
}

func (s *Comment) GetImages() string {
	return s.Image
}

func (s *Comment) GetText() string {
	return s.Text
}

func (s *Comment) GetName() string {
	return s.Name
}

func (s *Comment) GetAvatar() string {
	return s.Avatar
}

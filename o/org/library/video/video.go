package video

import "db/mgo"

type Video struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string   `json:"userid" bson:"userid"`
	BackGroundID  string   `json:"backgroundid" bson:"backgroundid"`
	Videos        []string `json:"videos" bson:"videos"`
}

func (s *Video) GetVideos() []string {
	return s.Videos
}

package photo

import "db/mgo"

type Photo struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string   `json:"userid" bson:"userid"`
	BackGroundID  string   `json:"backgroundid" bson:"backgroundid"`
	Images        []string `json:"images" bson:"images"`
}

func (s *Photo) GetImages() []string {
	return s.Images
}

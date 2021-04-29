package conversation

import "db/mgo"

type Conversation struct {
	mgo.BaseModel `bson:",inline"`
	UserIDs       []string `json:"userids" bson:"userids"`
	Images        []string `json:"images" bson:"images"`
	Files         []string `json:"files" bson:"files"`
	Links         []string `json:"links" bson:"links"`
}

func (s *Conversation) GetImages() []string {
	return s.Images
}

func (s *Conversation) GetFiles() []string {
	return s.Files
}

func (s *Conversation) GetLinks() []string {
	return s.Links
}

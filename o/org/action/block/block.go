package block

import "db/mgo"

type Block struct {
	mgo.BaseModel `bson:",inline"`
	Data          Content `json:"data" bson:"data"`
	UserID        string  `json:"userid" bson:"userid"`
	PersonID      string  `json:"personid" bson:"personid"`
}

type Content struct {
	Job     string `json:"job" bson:"job"`
	Content string `json:"content" bson:"content"`
}

func (s *Block) GetData() Content {
	return s.Data
}

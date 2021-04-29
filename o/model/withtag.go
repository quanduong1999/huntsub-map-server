package model

import (
	"db/mgo"
)

type WithTag struct {
	mgo.BaseModel `bson:",inline"`
	Type          string `bson:"type" json:"type"`
	Tag           string `bson:"tag" json:"tag"`
}

type IWithTag interface {
	mgo.IModel
}

func (t *TableWithType) GetByType(types []string, ptr interface{}) error {
	return t.ReadManyIn("type", types, ptr)
}

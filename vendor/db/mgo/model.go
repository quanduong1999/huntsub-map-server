package mgo

import (
	"time"
)

type IModel interface {
	GetID() string
	SetID(id string)
	IsDeleted() bool
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
}

type BaseModel struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`      //
	MTime int64  `bson:"mtime,omitempty" json:"mtime,omitempty"` //Modify Time
	DTime int64  `bson:"dtime" json:"dtime"`                     //Delete Time
}

func (b *BaseModel) GetID() string {
	return b.ID
}

func (b *BaseModel) SetID(s string) {
	b.ID = s
}

func (b *BaseModel) IsDeleted() bool {
	return b.DTime > 0
}

func (b *BaseModel) BeforeCreate() {
	b.MTime = time.Now().Unix()
	b.DTime = 0
}

func (b *BaseModel) BeforeUpdate() {
	b.MTime = time.Now().Unix()
}

func (b *BaseModel) BeforeDelete() {

}

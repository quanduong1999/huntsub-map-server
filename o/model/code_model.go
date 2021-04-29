package model

import (
	"db/mgo"
)

type WithCode struct {
	mgo.BaseModel `bson:",inline"`
	Code          string `bson:"code" json:"code"`
}

func (v *WithCode) GetCode() string {
	return v.Code
}

type IWithCode interface {
	mgo.IModel
	GetCode() string
}

func (t *TableWithCode) Create(v IWithCode) error {
	// check code
	err := t.NotExist(map[string]interface{}{"code": v.GetCode(), "dtime": 0})
	if err != nil {
		return err
	}
	return t.Table.Create(v)
}

func (t *TableWithCode) UpdateByID(oldID string, oldCode string, v IWithCode) error {
	if v.GetCode() != oldCode {
		// check code
		err := t.NotExist(map[string]interface{}{"code": v.GetCode(), "dtime": 0})
		if err != nil {
			return err
		}
	}
	return t.Table.UpdateByID(oldID, v)
}
func (t *TableWithCode) UpdateByIDIgnoreCheck(oldID string, oldCode string, v IWithCode) error {
	return t.Table.UpdateByID(oldID, v)
}

package mgo

type UnsafeModel interface {
	GetID() string
	SetID(id string)
}

//UnsafeBaseModel :
type UnsafeBaseModel struct {
	ID string `bson:"_id" json:"id"`
}

func (idm *UnsafeBaseModel) GetID() string {
	return idm.ID
}

func (idm *UnsafeBaseModel) SetID(id string) {
	idm.ID = id
}

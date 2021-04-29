package model

import (
	"db/mgo"
	"huntsub/huntsub-map-server/x/math"
)

type TableID struct {
	*mgo.UnsafeTable
}

func NewTableID(dbName, name string, idPrefix string) *mgo.UnsafeTable {
	var db = GetDB(dbName)
	var idMaker = math.RandStringMaker{Prefix: idPrefix, Length: 20}
	return mgo.NewUnsafeTable(db, name, &idMaker)
}

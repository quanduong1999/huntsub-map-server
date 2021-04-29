package model

import (
	"db/mgo"
	"huntsub/huntsub-map-server/config/cons"
	"huntsub/huntsub-map-server/x/math"
)

type TableWithBranchCode struct {
	*mgo.Table
}

type TableWithCode struct {
	*mgo.Table
}

type TableWithType struct {
	*mgo.Table
}

func NewTable(dbName, name string, idPrefix string) *mgo.Table {
	var db = GetDB(dbName)
	var idMaker = math.RandStringMaker{Prefix: idPrefix, Length: 20}
	return mgo.NewTable(db, name, &idMaker)
}

func NewTableWithCode(dbName, name string, idPrefix string) *TableWithCode {
	var table = NewTable(dbName, name, idPrefix)
	return &TableWithCode{Table: table}
}

func NewTableWithBranchCode(dbName, name string, idPrefix string) *TableWithBranchCode {
	var table = NewTable(dbName, name, idPrefix)
	return &TableWithBranchCode{Table: table}
}

func NewTableWithType(dbName, name string, idPrefix string) *TableWithType {
	var table = NewTable(dbName, name, idPrefix)
	return &TableWithType{Table: table}
}

func GetDB(dbName string) *mgo.Database {
	// if dbName == "did-home" {
	return mgo.GetDB(cons.ENV_HOME_DB)
	// } else if dbName == "did-asset" {
	// return mgo.GetDB(cons.ENV_LUCKY_DB)
	// } else {
	// 	return nil
	// }
}

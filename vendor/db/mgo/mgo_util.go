package mgo

import (
	"huntsub/huntsub-map-server/x/mlog"
)

var mongoDBLog = mlog.NewTagLog("MongoDB")

type M map[string]interface{}

package mgo

import (
	"http/web"
)

const errRecordNotFound = web.NotFound("record not found")
const errReadDataFailed = web.InternalServerError("read data failed")
const errInsertDataFailed = web.InternalServerError("insert data failed")
const errUpdateDataFailed = web.InternalServerError("update data failed")
const errRemoveDataFailed = web.InternalServerError("remove data failed")
const errCountDataFailed = web.InternalServerError("count data failed")
const errNoOutput = web.InternalServerError("no ouput for data")
const errDBClosed = web.InternalServerError("no connection to db")
const errMissingDbRegistration = web.InternalServerError("missing db registration")

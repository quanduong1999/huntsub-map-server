package database

import (
	"db/mgo"
	"fmt"
	"huntsub/huntsub-map-server/config/cons"
)

type DatabaseConfig struct {
	DBHost   string
	DBName   string
	Account  string
	Password string
}

func (o DatabaseConfig) String() string {
	return fmt.Sprintf("db:host=%s;name=%s", o.DBHost, o.DBName)
}

func (o *DatabaseConfig) Check() {
	mgo.Register(cons.ENV_HOME_DB, o.DBName, o.DBHost)
	// mgo.Register(cons.ENV_LUCKY_DB, "huntsub-map-server", o.DBHost)
}

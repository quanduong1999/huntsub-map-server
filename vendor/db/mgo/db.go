package mgo

import (
	"context"
	"mrw/event"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const dialTimeout = 10 * time.Second

//Database is mongo db
type Database struct {
	url  string
	name string
	*mgo.Database
	err        error
	connected  *event.Hub
	registered bool
	indexes    map[string][]mgo.Index
}

type DatabaseManager struct {
	databases map[string]*Database
	ping      chan *Database
	m         sync.Mutex
}

func newDbManager() *DatabaseManager {
	return &DatabaseManager{
		databases: map[string]*Database{},
		ping:      make(chan *Database, event.MediumHub),
	}
}

var dbm = newDbManager()

func GetDB(dbID string) *Database {
	dbm.m.Lock()
	defer dbm.m.Unlock()
	db := dbm.databases[dbID]
	if db == nil {
		db = &Database{
			err:       errDBClosed,
			connected: event.NewHub(event.SmallHub),
			indexes:   map[string][]mgo.Index{},
		}
		dbm.databases[dbID] = db
	}
	return db
}

func (db *Database) register(name string, url string) *Database {
	if db.registered {
		panic("db " + url + " already registered")
	}
	db.name = name
	db.url = url
	db.registered = true
	dbm.ping <- db
	return db
}

//NewDB connect to a new db
func Register(dbID, dbName string, url string) *Database {
	return GetDB(dbID).register(dbName, url)
}

func (db *Database) OnConnected() (event.Line, event.Cancel) {
	return db.connected.NewLine()
}

func (db *Database) reconnect() {
	session, err := mgo.DialWithTimeout(db.url, dialTimeout)
	db.err = err
	if err != nil {
		mongoDBLog.Errorf("connect to %v error %v", db.url, err.Error())
	} else {
		mongoDBLog.Infof(0, "connected to %v", db.url)
		session.SetSocketTimeout(time.Minute * 2)
		db.Database = session.DB(db.name)
		db.connected.Emit(struct{}{})
		for col, indecies := range db.indexes {
			for _, index := range indecies {
				db.C(col).EnsureIndex(index)
			}
		}
	}
}

func (db *Database) keepAlive() {
	if db == nil || !db.registered {
		return
	}
	if db.err != nil {
		db.reconnect()
		return
	}
	err := db.Session.Ping()
	if err != nil {
		mongoDBLog.Errorf("disconnected from %s", db.url)
		db.err = errDBClosed
	} else {
		mongoDBLog.Debugf(0, "ping %s success", db.url)
	}
}

func keepAlive() {
	for _, db := range dbm.databases {
		db.keepAlive()
	}
}

var startOnce sync.Once

func globalKeepAlive(ctx context.Context) {
	everyMinute := time.Tick(time.Minute)
	keepAlive()
	aliveContext, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-everyMinute:
			keepAlive()
		case db := <-dbm.ping:
			db.keepAlive()
		case <-aliveContext.Done():
			return
		}
	}
}

func Start(ctx context.Context) {
	startOnce.Do(func() {
		go globalKeepAlive(ctx)
	})
}

func (db *Database) ensureIndex(col string, index mgo.Index) {
	indecies, ok := db.indexes[col]
	if !ok {
		indecies = []mgo.Index{}
	}
	indecies = append(indecies, index)
	db.indexes[col] = indecies
}

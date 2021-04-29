package mgo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Index mgo.Index

type UnsafeTable struct {
	db   *Database
	Name string
	IdMaker
}

func NewUnsafeTable(db *Database, Name string, IDMake IdMaker) *UnsafeTable {
	return &UnsafeTable{
		db:      db,
		Name:    Name,
		IdMaker: IDMake,
	}
}

func (t *UnsafeTable) C() *mgo.Collection {
	return t.db.C(t.Name)
}

func (t *UnsafeTable) Col() (*mgo.Collection, error) {
	if t.db.err != nil {
		return nil, t.db.err
	}
	return t.db.C(t.Name), nil
}

func (t *UnsafeTable) UnsafeCreate(idm UnsafeModel) error {
	if t.IdMaker != nil {
		idm.SetID(t.IdMaker.Next())
	}
	return t.UnsafeInsert(idm)
}

func (t *UnsafeTable) UnsafeRunGetAll(query interface{}, ptr interface{}) error {
	if ptr == nil {
		return errNoOutput
	}
	collection, err := t.Col()
	if err != nil {
		return err
	}
	err = collection.Find(query).All(ptr)
	if err != nil {
		mongoDBLog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeRunGetOne(query interface{}, ptr interface{}) error {
	if ptr == nil {
		return errNoOutput
	}

	collection, err := t.Col()
	if err != nil {
		return err
	}

	var cursor = collection.Find(query).Iter()
	err = cursor.Err()
	if err != nil {
		mongoDBLog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	defer cursor.Close()

	if cursor.Next(ptr) {
		return nil
	}
	return errRecordNotFound
}

func (t *UnsafeTable) UnsafeInsert(obj interface{}) error {
	collection, err := t.Col()
	if err != nil {
		return err
	}
	err = collection.Insert(obj)
	if err != nil {
		mongoDBLog.ErrorDepth(2, err)
		return errInsertDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeCount(where interface{}) (int, error) {
	collection, err := t.Col()
	if err != nil {
		return 0, err
	}
	count, err := collection.Find(where).Count()
	if err != nil {
		mongoDBLog.ErrorDepth(2, err)
		return 0, errCountDataFailed
	}
	return count, nil
}

func (t *UnsafeTable) UnsafeGetByID(id string, ptr interface{}) error {
	return t.UnsafeRunGetOne(bson.M{"_id": id}, ptr)
}

func (t *UnsafeTable) UnsafeUpdateByID(id string, data interface{}) error {
	collection, err := t.Col()
	if err != nil {
		return err
	}
	err = collection.UpdateId(id, bson.M{"$set": data})
	if err != nil {
		if err == mgo.ErrNotFound {
			return errRecordNotFound
		}
		mongoDBLog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeUpsertByID(id string, data interface{}) error {
	collection, err := t.Col()
	if err != nil {
		return err
	}
	_, err = collection.UpsertId(id, bson.M{"$set": data})
	if err != nil {
		mongoDBLog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeDeleteByID(id string) error {
	collection, err := t.Col()
	if err != nil {
		return err
	}
	err = collection.RemoveId(id)
	if err != nil {
		mongoDBLog.ErrorDepth(2, err)
		return errRemoveDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeUpdateWhere(where, data interface{}) error {
	err := t.C().Update(where, bson.M{"$set": data})
	if err != nil {
		mongoDBLog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeReadAll(ptr interface{}) error {
	return t.UnsafeRunGetAll(nil, ptr)
}

func (t *UnsafeTable) UnsafeReadMany(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetAll(where, ptr)
}

func (t *UnsafeTable) UnsafeReadOne(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetOne(where, ptr)
}

func (t *UnsafeTable) EnsureIndex(index Index) int {
	t.db.ensureIndex(t.Name, mgo.Index(index))
	return 0
}

func (t *UnsafeTable) IsErrNotFound(err error) bool {
	return err == errRecordNotFound
}

func IsErrNotFound(err error) bool {
	return err == mgo.ErrNotFound || err == errRecordNotFound
}

func (t *UnsafeTable) EnsureAddressIndex() error {
	collection, err := t.Col()
	if err != nil {
		return err
	}
	// Might be needed one day
	pIndex := mgo.Index{
		Key:  []string{"$2dsphere:location"},
		Bits: 26,
	}
	err = collection.EnsureIndex(pIndex)
	if err != nil {
		return err
	}
	return nil
}

package channel

import "gopkg.in/mgo.v2/bson"

func GetChannel(where map[string]interface{}) (*Channel, error) {
	var u Channel
	return &u, TableChannel.ReadOne(where, &u)
}

func GetByID(id string) (*Channel, error) {
	var u Channel
	return &u, TableChannel.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Channel, error) {
	var b = []*Channel{}
	return b, TableChannel.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Channel, error) {
	var Channels = []*Channel{}
	return Channels, TableChannel.UnsafeReadAll(&Channels)
}

func GetSeach(key, value string) ([]*Channel, error) {
	var Channels = []*Channel{}
	return Channels, TableChannel.Search(key, value, &Channels)
}

func Count(where map[string]interface{}) (int, error) {
	return TableChannel.UnsafeCount(where)
}

func GetPaginationTop(where map[string]interface{}, skip, limit int) ([]*Channel, error) {
	var rans = []*Channel{}
	var sort = bson.M{"mtime": -1}
	pipe := []bson.M{
		{"$match": where},
		{"$sort": sort},
		{"$skip": (skip - 1) * limit},
		{"$limit": limit},
	}
	return rans, TableChannel.C().Pipe(pipe).All(&rans)
}

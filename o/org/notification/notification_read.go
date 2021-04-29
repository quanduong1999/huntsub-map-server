package notification

import "gopkg.in/mgo.v2/bson"

func GetNotification(where map[string]interface{}) (*NotifiationManagement, error) {
	var u NotifiationManagement
	return &u, TableNotifiationManagement.ReadOne(where, &u)
}

func GetByID(id string) (*NotifiationManagement, error) {
	var u NotifiationManagement
	return &u, TableNotifiationManagement.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*NotifiationManagement, error) {
	var b = []*NotifiationManagement{}
	return b, TableNotifiationManagement.UnsafeReadMany(where, &b)
}

func GetAll() ([]*NotifiationManagement, error) {
	var Posts = []*NotifiationManagement{}
	return Posts, TableNotifiationManagement.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*NotifiationManagement, error) {
	var Posts = []*NotifiationManagement{}
	return Posts, TableNotifiationManagement.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TableNotifiationManagement.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*NotifiationManagement, error) {
	var tks = []*NotifiationManagement{}
	var err error
	if _type != "ascending" {
		err = TableNotifiationManagement.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableNotifiationManagement.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

func GetPaginationTop(where map[string]interface{}, skip, limit int) ([]*NotifiationManagement, error) {
	var rans = []*NotifiationManagement{}
	var sort = bson.M{"experiencenumber": -1}
	pipe := []bson.M{
		{"$match": where},
		{"$sort": sort},
		{"$skip": (skip - 1) * 10},
		{"$limit": limit},
	}
	return rans, TableNotifiationManagement.C().Pipe(pipe).All(&rans)
}

package rank

import "gopkg.in/mgo.v2/bson"

func GetRank(where map[string]interface{}) (*Rank, error) {
	var u Rank
	return &u, TableRank.ReadOne(where, &u)
}

func GetByID(id string) (*Rank, error) {
	var u Rank
	return &u, TableRank.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Rank, error) {
	var b = []*Rank{}
	return b, TableRank.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Rank, error) {
	var Posts = []*Rank{}
	return Posts, TableRank.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*Rank, error) {
	var Posts = []*Rank{}
	return Posts, TableRank.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TableRank.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Rank, error) {
	var tks = []*Rank{}
	var err error
	if _type != "ascending" {
		err = TableRank.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableRank.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

func GetPaginationTop(where map[string]interface{}, skip, limit int) ([]*Rank, error) {
	var rans = []*Rank{}
	var sort = bson.M{"experiencenumber": -1}
	pipe := []bson.M{
		{"$match": where},
		{"$sort": sort},
		{"$skip": (skip - 1) * 10},
		{"$limit": limit},
	}
	return rans, TableRank.C().Pipe(pipe).All(&rans)
}

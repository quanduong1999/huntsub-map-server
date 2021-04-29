package post

import "gopkg.in/mgo.v2/bson"

func GetPost(where map[string]interface{}) (*Post, error) {
	var u Post
	return &u, TablePost.ReadOne(where, &u)
}

func GetByID(id string) (*Post, error) {
	var u Post
	return &u, TablePost.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Post, error) {
	var b = []*Post{}
	return b, TablePost.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Post, error) {
	var Posts = []*Post{}
	return Posts, TablePost.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*Post, error) {
	var Posts = []*Post{}
	return Posts, TablePost.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TablePost.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Post, error) {
	var tks = []*Post{}
	var err error
	if _type != "ascending" {
		err = TablePost.C().Find(where).Sort("$mtime").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TablePost.C().Find(where).Sort("mtime", "").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

func GetPostNearMe(where map[string]interface{}, skip, limit int) ([]*Post, error) {
	var posts = []*Post{}
	var sort = bson.M{"mtime": -1}
	pipe := []bson.M{
		{"$match": where},
		{"$sort": sort},
		{"$skip": (skip - 1) * 10},
		{"$limit": limit},
	}
	return posts, TablePost.C().Pipe(pipe).All(&posts)
}

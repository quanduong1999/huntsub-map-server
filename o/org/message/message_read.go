package message

func GetOne(where map[string]interface{}) (*Message, error) {
	var u Message
	return &u, TableMessage.ReadOne(where, &u)
}

func GetByID(id string) (*Message, error) {
	var u Message
	return &u, TableMessage.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Message, error) {
	var b = []*Message{}
	return b, TableMessage.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Message, error) {
	var Posts = []*Message{}
	return Posts, TableMessage.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*Message, error) {
	var Posts = []*Message{}
	return Posts, TableMessage.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TableMessage.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Message, error) {
	var tks = []*Message{}
	var err error
	if _type != "ascending" {
		err = TableMessage.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableMessage.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

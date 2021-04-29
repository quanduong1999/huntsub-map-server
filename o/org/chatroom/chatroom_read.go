package chatroom

func GetOne(where map[string]interface{}) (*ChatRoom, error) {
	var u ChatRoom
	return &u, TableChatRoom.ReadOne(where, &u)
}

func GetByID(id string) (*ChatRoom, error) {
	var u ChatRoom
	return &u, TableChatRoom.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*ChatRoom, error) {
	var b = []*ChatRoom{}
	return b, TableChatRoom.UnsafeReadMany(where, &b)
}

func GetAll() ([]*ChatRoom, error) {
	var Posts = []*ChatRoom{}
	return Posts, TableChatRoom.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*ChatRoom, error) {
	var Posts = []*ChatRoom{}
	return Posts, TableChatRoom.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TableChatRoom.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*ChatRoom, error) {
	var tks = []*ChatRoom{}
	var err error
	if _type != "ascending" {
		err = TableChatRoom.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableChatRoom.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

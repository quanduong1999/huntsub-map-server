package user

func GetWHFeedBack(where map[string]interface{}) (*UserData, error) {
	var u UserData
	return &u, TableUserData.ReadOne(where, &u)
}

func GetByID(id string) (*UserData, error) {
	var u UserData
	return &u, TableUserData.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*UserData, error) {
	var b = []*UserData{}
	return b, TableUserData.UnsafeReadMany(where, &b)
}

func GetAll() ([]*UserData, error) {
	var WHFeedBacks = []*UserData{}
	return WHFeedBacks, TableUserData.UnsafeReadAll(&WHFeedBacks)
}

func GetSeach(key, value string) ([]*UserData, error) {
	var WHFeedBacks = []*UserData{}
	return WHFeedBacks, TableUserData.Search(key, value, &WHFeedBacks)
}

func Count(where map[string]interface{}) (int, error) {
	return TableUserData.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*UserData, error) {
	var tks = []*UserData{}
	var err error
	if _type != "ascending" {
		err = TableUserData.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableUserData.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

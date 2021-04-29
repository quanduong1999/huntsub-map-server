package whfeedback

func GetWHFeedBack(where map[string]interface{}) (*WHFeedback, error) {
	var u WHFeedback
	return &u, TableWHFeedBack.ReadOne(where, &u)
}

func GetByID(id string) (*WHFeedback, error) {
	var u WHFeedback
	return &u, TableWHFeedBack.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*WHFeedback, error) {
	var b = []*WHFeedback{}
	return b, TableWHFeedBack.UnsafeReadMany(where, &b)
}

func GetAll() ([]*WHFeedback, error) {
	var WHFeedBacks = []*WHFeedback{}
	return WHFeedBacks, TableWHFeedBack.UnsafeReadAll(&WHFeedBacks)
}

func GetSeach(key, value string) ([]*WHFeedback, error) {
	var WHFeedBacks = []*WHFeedback{}
	return WHFeedBacks, TableWHFeedBack.Search(key, value, &WHFeedBacks)
}

func Count(where map[string]interface{}) (int, error) {
	return TableWHFeedBack.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*WHFeedback, error) {
	var tks = []*WHFeedback{}
	var err error
	if _type != "ascending" {
		err = TableWHFeedBack.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableWHFeedBack.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

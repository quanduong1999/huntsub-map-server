package feedback

func getFeedBack(where map[string]interface{}) (*Feedback, error) {
	var u Feedback
	return &u, TableFeedBack.ReadOne(where, &u)
}

func GetByID(id string) (*Feedback, error) {
	var u Feedback
	return &u, TableFeedBack.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Feedback, error) {
	var b = []*Feedback{}
	return b, TableFeedBack.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Feedback, error) {
	var FeedBacks = []*Feedback{}
	return FeedBacks, TableFeedBack.UnsafeReadAll(&FeedBacks)
}

func GetSeach(key, value string) ([]*Feedback, error) {
	var FeedBacks = []*Feedback{}
	return FeedBacks, TableFeedBack.Search(key, value, &FeedBacks)
}

func Count(where map[string]interface{}) (int, error) {
	return TableFeedBack.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Feedback, error) {
	var tks = []*Feedback{}
	var err error
	if _type != "ascending" {
		err = TableFeedBack.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableFeedBack.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

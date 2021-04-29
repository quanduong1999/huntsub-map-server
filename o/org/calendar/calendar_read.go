package calendar

func GetPost(where map[string]interface{}) (*Calendar, error) {
	var u Calendar
	return &u, TableCalendar.ReadOne(where, &u)
}

func GetByID(id string) (*Calendar, error) {
	var u Calendar
	return &u, TableCalendar.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Calendar, error) {
	var b = []*Calendar{}
	return b, TableCalendar.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Calendar, error) {
	var Posts = []*Calendar{}
	return Posts, TableCalendar.UnsafeReadAll(&Posts)
}

func GetSeach(key, value string) ([]*Calendar, error) {
	var Posts = []*Calendar{}
	return Posts, TableCalendar.Search(key, value, &Posts)
}

func Count(where map[string]interface{}) (int, error) {
	return TableCalendar.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Calendar, error) {
	var tks = []*Calendar{}
	var err error
	if _type != "ascending" {
		err = TableCalendar.C().Find(where).Sort("-meeting_time").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableCalendar.C().Find(where).Sort("meeting_time").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

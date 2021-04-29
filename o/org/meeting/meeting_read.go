package meeting

func getMeeting(where map[string]interface{}) (*Meeting, error) {
	var u Meeting
	return &u, TableMeeting.ReadOne(where, &u)
}

func GetByID(id string) (*Meeting, error) {
	var u Meeting
	return &u, TableMeeting.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Meeting, error) {
	var b = []*Meeting{}
	return b, TableMeeting.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Meeting, error) {
	var Meetings = []*Meeting{}
	return Meetings, TableMeeting.UnsafeReadAll(&Meetings)
}

func GetSeach(key, value string) ([]*Meeting, error) {
	var Meetings = []*Meeting{}
	return Meetings, TableMeeting.Search(key, value, &Meetings)
}

func Count(where map[string]interface{}) (int, error) {
	return TableMeeting.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Meeting, error) {
	var tks = []*Meeting{}
	var err error
	if _type != "ascending" {
		err = TableMeeting.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableMeeting.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

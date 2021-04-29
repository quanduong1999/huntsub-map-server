package user

func GetUserReport(where map[string]interface{}) (*UserReport, error) {
	var u UserReport
	return &u, TableUserReport.ReadOne(where, &u)
}

func GetByID(id string) (*UserReport, error) {
	var u UserReport
	return &u, TableUserReport.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*UserReport, error) {
	var b = []*UserReport{}
	return b, TableUserReport.UnsafeReadMany(where, &b)
}

func GetAll() ([]*UserReport, error) {
	var UserReports = []*UserReport{}
	return UserReports, TableUserReport.UnsafeReadAll(&UserReports)
}

func GetSeach(key, value string) ([]*UserReport, error) {
	var UserReports = []*UserReport{}
	return UserReports, TableUserReport.Search(key, value, &UserReports)
}

func Count(where map[string]interface{}) (int, error) {
	return TableUserReport.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*UserReport, error) {
	var tks = []*UserReport{}
	var err error
	if _type != "ascending" {
		err = TableUserReport.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableUserReport.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

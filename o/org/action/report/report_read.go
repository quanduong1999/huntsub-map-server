package report

func getReport(where map[string]interface{}) (*Report, error) {
	var u Report
	return &u, TableReport.ReadOne(where, &u)
}

func GetByID(id string) (*Report, error) {
	var u Report
	return &u, TableReport.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Report, error) {
	var b = []*Report{}
	return b, TableReport.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Report, error) {
	var Reports = []*Report{}
	return Reports, TableReport.UnsafeReadAll(&Reports)
}

func GetSeach(key, value string) ([]*Report, error) {
	var Reports = []*Report{}
	return Reports, TableReport.Search(key, value, &Reports)
}

func Count(where map[string]interface{}) (int, error) {
	return TableReport.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Report, error) {
	var tks = []*Report{}
	var err error
	if _type != "ascending" {
		err = TableReport.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableReport.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

package share

func getShare(where map[string]interface{}) (*Share, error) {
	var u Share
	return &u, TableShare.ReadOne(where, &u)
}

func GetByID(id string) (*Share, error) {
	var u Share
	return &u, TableShare.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Share, error) {
	var b = []*Share{}
	return b, TableShare.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Share, error) {
	var Shares = []*Share{}
	return Shares, TableShare.UnsafeReadAll(&Shares)
}

func GetSeach(key, value string) ([]*Share, error) {
	var Shares = []*Share{}
	return Shares, TableShare.Search(key, value, &Shares)
}

func Count(where map[string]interface{}) (int, error) {
	return TableShare.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Share, error) {
	var tks = []*Share{}
	var err error
	if _type != "ascending" {
		err = TableShare.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableShare.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

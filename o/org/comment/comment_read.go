package comment

func getComment(where map[string]interface{}) (*Comment, error) {
	var u Comment
	return &u, TableComment.ReadOne(where, &u)
}

func GetByID(id string) (*Comment, error) {
	var u Comment
	return &u, TableComment.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Comment, error) {
	var b = []*Comment{}
	return b, TableComment.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Comment, error) {
	var Comments = []*Comment{}
	return Comments, TableComment.UnsafeReadAll(&Comments)
}

func GetSeach(key, value string) ([]*Comment, error) {
	var Comments = []*Comment{}
	return Comments, TableComment.Search(key, value, &Comments)
}

func Count(where map[string]interface{}) (int, error) {
	return TableComment.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Comment, error) {
	var tks = []*Comment{}
	var err error
	if _type != "ascending" {
		err = TableComment.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableComment.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

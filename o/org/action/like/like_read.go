package like

func GetLike(where map[string]interface{}) (*Like, error) {
	var u Like
	return &u, TableLike.ReadOne(where, &u)
}

func GetByID(id string) (*Like, error) {
	var u Like
	return &u, TableLike.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Like, error) {
	var b = []*Like{}
	return b, TableLike.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Like, error) {
	var Likes = []*Like{}
	return Likes, TableLike.UnsafeReadAll(&Likes)
}

func GetSeach(key, value string) ([]*Like, error) {
	var Likes = []*Like{}
	return Likes, TableLike.Search(key, value, &Likes)
}

func Count(where map[string]interface{}) (int, error) {
	return TableLike.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Like, error) {
	var tks = []*Like{}
	var err error
	if _type != "ascending" {
		err = TableLike.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableLike.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

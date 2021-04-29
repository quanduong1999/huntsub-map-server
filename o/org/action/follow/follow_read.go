package follow

func GetFollow(where map[string]interface{}) (*Follow, error) {
	var u Follow
	return &u, TableFollow.ReadOne(where, &u)
}

func GetByID(id string) (*Follow, error) {
	var u Follow
	return &u, TableFollow.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Follow, error) {
	var b = []*Follow{}
	return b, TableFollow.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Follow, error) {
	var Follows = []*Follow{}
	return Follows, TableFollow.UnsafeReadAll(&Follows)
}

func GetSeach(key, value string) ([]*Follow, error) {
	var Follows = []*Follow{}
	return Follows, TableFollow.Search(key, value, &Follows)
}

func Count(where map[string]interface{}) (int, error) {
	return TableFollow.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Follow, error) {
	var tks = []*Follow{}
	var err error
	if _type != "ascending" {
		err = TableFollow.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableFollow.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

package photo

func getPhoto(where map[string]interface{}) (*Photo, error) {
	var u Photo
	return &u, TablePhoto.ReadOne(where, &u)
}

func GetByID(id string) (*Photo, error) {
	var u Photo
	return &u, TablePhoto.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Photo, error) {
	var b = []*Photo{}
	return b, TablePhoto.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Photo, error) {
	var Photos = []*Photo{}
	return Photos, TablePhoto.UnsafeReadAll(&Photos)
}

func GetSeach(key, value string) ([]*Photo, error) {
	var Photos = []*Photo{}
	return Photos, TablePhoto.Search(key, value, &Photos)
}

func Count(where map[string]interface{}) (int, error) {
	return TablePhoto.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Photo, error) {
	var tks = []*Photo{}
	var err error
	if _type != "ascending" {
		err = TablePhoto.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TablePhoto.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

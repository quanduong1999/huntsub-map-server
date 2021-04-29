package video

func getVideo(where map[string]interface{}) (*Video, error) {
	var u Video
	return &u, TableVideo.ReadOne(where, &u)
}

func GetByID(id string) (*Video, error) {
	var u Video
	return &u, TableVideo.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Video, error) {
	var b = []*Video{}
	return b, TableVideo.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Video, error) {
	var Videos = []*Video{}
	return Videos, TableVideo.UnsafeReadAll(&Videos)
}

func GetSeach(key, value string) ([]*Video, error) {
	var Videos = []*Video{}
	return Videos, TableVideo.Search(key, value, &Videos)
}

func Count(where map[string]interface{}) (int, error) {
	return TableVideo.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Video, error) {
	var tks = []*Video{}
	var err error
	if _type != "ascending" {
		err = TableVideo.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableVideo.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

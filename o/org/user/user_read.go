package user

func getUser(where map[string]interface{}) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(where, &u)
}

func GetByID(id string) (*User, error) {
	var u User
	return &u, TableUser.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*User, error) {
	var b = []*User{}
	return b, TableUser.UnsafeReadMany(where, &b)
}

func GetByUsername(username string) (*User, error) {
	var u User
	return &u, TableUser.ReadOne(map[string]interface{}{
		"username": username,
		"dtime":    0,
	}, &u)
}

func CheckUsernamePassword(username string, password string) (bool, error) {
	u, err := GetByUsername(username)
	if err = u.ComparePassword(password); err != nil {
		return false, err
	}
	return true, err
}

func GetAll() ([]*User, error) {
	var users = []*User{}
	return users, TableUser.UnsafeReadAll(&users)
}

func GetSeach(key, value string) ([]*User, error) {
	var users = []*User{}
	return users, TableUser.Search(key, value, &users)
}

func Count(where map[string]interface{}) (int, error) {
	return TableUser.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*User, error) {
	var tks = []*User{}
	var err error
	if _type != "ascending" {
		err = TableUser.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableUser.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}

func GetByPagination(where map[string]interface{}, skip, limit int) ([]*User, error) {
	var tks = []*User{}
	var err error
	err = TableUser.C().Find(where).Skip((skip - 1) * limit).Limit(limit).All(&tks)

	return tks, err
}

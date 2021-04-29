package user

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/user"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheUserLog = mlog.NewTagLog("cache_user")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	usr, _ := o.Data.(*user.User)
	cacheUserLog.Infof(0, "Handle Event User %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return Set(usr.GetID(), usr)
	case oev.ObjectActionUpdate:
		return Update(usr.GetID(), usr)
	case oev.ObjectActionMarkDelete:
		return Del(usr.GetID())
	}
	return nil
}

func Set(k string, v *user.User) error {
	val, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var redisdb = redis.GetRedisClient()
	err = redisdb.Set(k, val, redis.Exprired).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(k string) (*user.User, error) {
	var c = &user.User{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := user.GetByID(k)
		if err != nil {
			return nil, err
		}
		Set(k, com)
		return com, nil
	}
	err = json.Unmarshal([]byte(val), c)
	if err != nil {
		return nil, err
	}
	Set(k, c)
	return c, nil
}

func Del(k string) error {
	var redisdb = redis.GetRedisClient()
	err := redisdb.Del(k).Err()
	return err
}

func Update(k string, v *user.User) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

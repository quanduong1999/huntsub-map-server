package follow

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/follow"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheFollowLog = mlog.NewTagLog("cache_follow")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	cal, _ := o.Data.(*follow.Follow)
	cacheFollowLog.Infof(0, "Handle Event Follow %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return Set(cal.GetID(), cal)
	case oev.ObjectActionUpdate:
		return Update(cal.GetID(), cal)
	case oev.ObjectActionMarkDelete:
		return Del(cal.GetID())
	}
	return nil
}

func Set(k string, v *follow.Follow) error {
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

func Get(k string) (*follow.Follow, error) {
	var c = &follow.Follow{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := follow.GetByID(k)
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

func Update(k string, v *follow.Follow) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

package share

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/action/share"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheShareLog = mlog.NewTagLog("cache_share")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	cal, _ := o.Data.(*share.Share)
	cacheShareLog.Infof(0, "Handle Event Share %s", o.Action)
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

func Set(k string, v *share.Share) error {
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

func Get(k string) (*share.Share, error) {
	var c = &share.Share{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := share.GetByID(k)
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

func Update(k string, v *share.Share) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

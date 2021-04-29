package whfeedback

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheWHFeedbackLog = mlog.NewTagLog("cache_whfeedback")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	pos, _ := o.Data.(*whfeedback.WHFeedback)
	cacheWHFeedbackLog.Infof(0, "Handle Event WHFeedback %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return Set(pos.GetID(), pos)
	case oev.ObjectActionUpdate:
		return Update(pos.GetID(), pos)
	case oev.ObjectActionMarkDelete:
		return Del(pos.GetID())
	}
	return nil
}

func Set(k string, v *whfeedback.WHFeedback) error {
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

func Get(k string) (*whfeedback.WHFeedback, error) {
	var c = &whfeedback.WHFeedback{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := whfeedback.GetByID(k)
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

func Update(k string, v *whfeedback.WHFeedback) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

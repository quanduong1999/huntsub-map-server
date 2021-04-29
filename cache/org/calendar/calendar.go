package calendar

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/calendar"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheCalendarLog = mlog.NewTagLog("cache_post")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	cal, _ := o.Data.(*calendar.Calendar)
	cacheCalendarLog.Infof(0, "Handle Event Calendar %s", o.Action)
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

func Set(k string, v *calendar.Calendar) error {
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

func Get(k string) (*calendar.Calendar, error) {
	var c = &calendar.Calendar{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := calendar.GetByID(k)
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

func Update(k string, v *calendar.Calendar) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

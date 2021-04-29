package rank

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheRankLog = mlog.NewTagLog("cache_rank")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	ran, _ := o.Data.(*rank.Rank)
	cacheRankLog.Infof(0, "Handle Event Rank %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return Set(ran.GetID(), ran)
	case oev.ObjectActionUpdate:
		return Update(ran.GetID(), ran)
	case oev.ObjectActionMarkDelete:
		return Del(ran.GetID())
	}
	return nil
}

func Set(k string, v *rank.Rank) error {
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

func Get(k string) (*rank.Rank, error) {
	var c = &rank.Rank{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := rank.GetByID(k)
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

func Update(k string, v *rank.Rank) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

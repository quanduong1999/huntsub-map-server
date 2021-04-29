package channel

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cacheChannelLog = mlog.NewTagLog("cache_channel")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	cn, _ := o.Data.(*channel.Channel)
	cacheChannelLog.Infof(0, "Handle Event Channel %s", o.Action)
	switch o.Action {
	case oev.ObjectActionCreate:
		return Set(cn.GetID(), cn)
	case oev.ObjectActionUpdate:
		return Update(cn.GetID(), cn)
	case oev.ObjectActionMarkDelete:
		return Del(cn.GetID())
	}
	return nil
}

func Set(k string, v *channel.Channel) error {
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

func Get(k string) (*channel.Channel, error) {
	var c = &channel.Channel{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := channel.GetByID(k)
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

func Update(k string, v *channel.Channel) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

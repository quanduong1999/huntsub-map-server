package photo

import (
	"encoding/json"
	"huntsub/huntsub-map-server/event/shared/oev"
	"huntsub/huntsub-map-server/o/org/library/photo"
	"huntsub/huntsub-map-server/x/mlog"
	"huntsub/huntsub-map-server/x/redis"
	"util/runtime"
)

var cachePhotoLog = mlog.NewTagLog("cache_photo")

func Handle(o *oev.ObjectEvent) error {
	defer runtime.Recover()
	pos, _ := o.Data.(*photo.Photo)
	cachePhotoLog.Infof(0, "Handle Event Photo %s", o.Action)
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

func Set(k string, v *photo.Photo) error {
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

func Get(k string) (*photo.Photo, error) {
	var c = &photo.Photo{}
	var redisdb = redis.GetRedisClient()
	val, err := redisdb.Get(k).Result()
	if err != nil {
		com, err := photo.GetByID(k)
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

func Update(k string, v *photo.Photo) error {
	_ = Del(k)
	err := Set(k, v)
	return err
}

package redis

import (
	"huntsub/huntsub-map-server/x/mlog"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client
var Exprired = 300 * time.Second
var cacheRedisLog = mlog.NewTagLog("cache_redis")

func GetRedisClient() *redis.Client {
	return redisdb
}

func Start() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pong, err := redisdb.Ping().Result()
	err = redisdb.FlushAll().Err()
	cacheRedisLog.Infof(0, "redis is ready %s", pong, err)
}

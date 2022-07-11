package util

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var (
	SubPool *redis.Pool
)

// deprecated
//func init() {
//	redisHost := GetEnv("REDIS_HOST", "127.0.0.1")
//	redisPort := GetEnv("REDIS_PORT", "6379")
//	redisPassword := GetEnv("REDIS_PASSWORD", "")
//	redisDatabase := GetEnv("REDIS_DATABASE", "0")
//	SubPool = initRedis(redisDatabase, redisPassword, redisHost+":"+redisPort)
//}

func initRedis(database string, password string, server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			db, _ := strconv.Atoi(database)
			c, err := redis.Dial("tcp", server, redis.DialPassword(password), redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

package util

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var cacheKey = func(key string) string {
	return fmt.Sprintf("link:%s", key)
}

func SetCache(key string, value interface{}, ttl int) (err error) {
	conn := SubPool.Get()
	defer conn.Close()
	var strValue string
	switch value.(type) {
	case string:
		strValue = value.(string)
	case []byte:
		strValue = string(value.([]byte))
	case int:
		strValue = IntToString(value.(int))
	default:
		b, _ := json.Marshal(value)
		strValue = string(b)
	}
	_, err = conn.Do("setex", cacheKey(key), ttl, strValue)
	return err
}

func GetCache(key string) []byte {
	conn := SubPool.Get()
	defer conn.Close()
	if cache, err := redis.String(conn.Do("get", cacheKey(key))); err != nil {
		return nil
	} else {
		return []byte(cache)
	}
}
func GetCacheUint64(key string) uint64 {
	conn := SubPool.Get()
	defer conn.Close()
	if cache, err := redis.Uint64(conn.Do("get", cacheKey(key))); err != nil {
		return 0
	} else {
		return cache
	}
}


func DelCache(key string) {
	conn := SubPool.Get()
	defer conn.Close()
	conn.Do("del", cacheKey(key))
}

func SaddCache(key, value string) bool {
	conn := SubPool.Get()
	defer conn.Close()
	if intReturn, err := redis.Int(conn.Do("sadd", cacheKey(key), value)); err != nil || intReturn != 1 {
		return false
	} else {
		return true
	}
}

func SremCache(key, value string) bool {
	conn := SubPool.Get()
	defer conn.Close()
	if intReturn, err := redis.Int(conn.Do("srem", cacheKey(key), value)); err != nil || intReturn != 1 {
		return false
	} else {
		return true
	}
}

func SmembersCache(key string) []string {
	conn := SubPool.Get()
	defer conn.Close()
	intReturn, _ := redis.Strings(conn.Do("smembers", cacheKey(key)))
	return intReturn

}

func SaddArray(key string, value []interface{}) bool {
	conn := SubPool.Get()
	defer conn.Close()
	data := []interface{}{cacheKey(key)}
	value = append(data, value...)
	if intReturn, err := redis.Int(conn.Do("sadd", value...)); err != nil || intReturn < 1 {
		return false
	} else {
		return true
	}

}

func HgetCache(key, field string) []byte {
	conn := SubPool.Get()
	defer conn.Close()
	if cache, err := redis.String(conn.Do("hget", cacheKey(key), field)); err != nil {
		return nil
	} else {
		return []byte(cache)
	}
}

func HgetCacheAll(key string) map[string]string {
    conn := SubPool.Get()
    defer conn.Close()
    if cache, err := redis.Strings(conn.Do("hgetall", cacheKey(key))); err != nil {
        return nil
    } else {
        ret := make(map[string]string)
        for index := 0; index < len(cache); index += 2 {
            ret[cache[index]] = cache[index+1]
        }
        return ret
    }
}

func HsetCache(key, field string, value []byte) []byte {
	conn := SubPool.Get()
	defer conn.Close()
	if cache, err := redis.String(conn.Do("hset", cacheKey(key), field, string(value))); err != nil {
		return nil
	} else {
		return []byte(cache)
	}
}

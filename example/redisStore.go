package example

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/mojocn/ip2location"
	"log"
)

const RedisAddr = "127.0.0.1:6379"
const RedisPassword = "8796534shdjq384wejgkguiern"
const RedisDBIndex = 0

type RedisIpCache struct {
	redis          *redis.Client
	RedisIpHashKey string
}

func (c *RedisIpCache) GetLocation(ip string) (lo *ip2location.Location) {
	if data, err := c.redis.HGet(c.RedisIpHashKey, ip).Bytes(); err == nil {
		json.Unmarshal(data, lo)
		return
	} else {
		return nil
	}
}

func (c *RedisIpCache) SetLocation(ip string, lo *ip2location.Location) {
	if bytes, err := json.Marshal(lo); err == nil {
		c.redis.HSet(c.RedisIpHashKey, ip, bytes)
	} else {
		log.Println(err)
	}
}

func newIpCacheDriver() *RedisIpCache {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword, // no password set
		DB:       RedisDBIndex,  // use default DB
	})

	return &RedisIpCache{
		redis:          client,
		RedisIpHashKey: "myIpDb",
	}
}

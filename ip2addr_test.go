package ip2location

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"testing"
)

const RedisAddr = "127.0.0.1:6379"
const RedisPassword = "8796534shdjq384wejgkguiern"
const RedisDBIndex = 0

type redisIpCache struct {
	redis     *redis.Client
	IpHashKey string
}

func newIpCacheDriver() *redisIpCache {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword, // no password set
		DB:       RedisDBIndex,  // use default DB
	})

	return &redisIpCache{
		redis:     client,
		IpHashKey: "myIpDb",
	}
}

func (c *redisIpCache) GetLocation(ip string) (lo *Location) {
	if data, err := c.redis.HGet(c.IpHashKey, ip).Bytes(); err == nil {
		json.Unmarshal(data, lo)
		return
	} else {
		return nil
	}
}
func (c *redisIpCache) SetLocation(ip string, lo *Location) {
	if bytes, err := json.Marshal(lo); err == nil {
		c.redis.HSet(c.IpHashKey, ip, bytes)
	} else {
		log.Println(err)
	}
}

func TestIpParser_FetchIpAddress(t *testing.T) {

	parser := IpParser{
		Handlers: []locationFetcher{
			fetcherTaobao,
			fetcherSina,
			fetcherBaidu,
		},
		Store: newIpCacheDriver(),
	}

	parser.FetchIpAddress("119.96.211.173")

	loc, err := fetcherBaidu("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

	loc, err = fetcherSina("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

	loc, err = fetcherTaobao("119.96.211.173")
	if err != nil {
		t.Error(err)
	}
	log.Print(loc)

}

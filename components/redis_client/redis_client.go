package redis_client

import (
	"github.com/go-redis/redis"
	services "gocherry-api-gateway/proxy/service"
	"time"
)

type Client struct {
	baseClient *redis.Client
}

//内部调用
func RedisClient(class string) *Client {
	Addr := ""
	Password := ""
	DB := 0

	appConfig := services.GetAppConfig()

	switch class {
	case "proxy":
		Addr = appConfig.Common.RedisProxyIp + ":" + appConfig.Common.RedisProxyPort
		Password = appConfig.Common.RedisProxyPass
		DB = 0
		break
	default:
		Addr = appConfig.Common.RedisProxyIp + ":" + appConfig.Common.RedisProxyPort
		Password = appConfig.Common.RedisProxyPass
		DB = 0
		break
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
	client := &Client{
		baseClient: redisClient,
	}
	return client
}

//每新增一个实例就这么加一个获取实例的方法
func GetProxyRedis() *Client {
	return RedisClient("proxy")
}

func GetDefaultRedis() *Client {
	return RedisClient("default")
}

func (c *Client) Set(key string, value string, expiration time.Duration) interface{} {
	ret, _ := c.baseClient.Set(key, value, expiration).Result()
	return ret
}

func (c Client) Get(key string) string {
	ret, _ := c.baseClient.Get(key).Result()
	return ret
}

func (c Client) Del(key string) interface{} {
	ret, _ := c.baseClient.Del(key).Result()
	return ret
}

//hash

func (c Client) HSet(key string, field string, value interface{}, expiration time.Duration) interface{} {
	ret, _ := c.baseClient.HSet(key, field, value).Result()
	return ret
}

func (c Client) HGet(key string, field string) interface{} {
	ret, _ := c.baseClient.HGet(key, field).Result()
	return ret
}

func (c Client) HGetAll(key string) map[string]string {
	ret, _ := c.baseClient.HGetAll(key).Result()
	return ret
}

func (c Client) HLen(key string) int {
	ret, _ := c.baseClient.HLen(key).Result()
	return int(ret)
}

func (c Client) HDel(key string) interface{} {
	ret, _ := c.baseClient.HDel(key).Result()
	return ret
}

func (c Client) HExists(key string, field string) bool {
	ret, _ := c.baseClient.HExists(key, field).Result()
	return ret
}

func (c Client) HIncrBy(key string, field string, num int64, expiration time.Duration) interface{} {
	ret, _ := c.baseClient.HIncrBy(key, field, int64(num)).Result()
	return ret
}

func (c Client) HMget(key string, fields string) interface{} {
	ret, _ := c.baseClient.HMGet(key, fields).Result()
	return ret
}

func (c Client) HMset(key string, fields map[string]interface{}) interface{} {
	ret, _ := c.baseClient.HMSet(key, fields).Result()
	return ret
}

func (c Client) HVals(key string) interface{} {
	ret, _ := c.baseClient.HVals(key).Result()
	return ret
}

//list
func (c Client) HPush(key string, value interface{}) interface{} {
	ret, _ := c.baseClient.LPush(key, value).Result()
	return ret
}

func (c Client) HPop(key string) interface{} {
	ret, _ := c.baseClient.RPop(key).Result()
	return ret
}

func (c Client) LPush(key string, value interface{}) interface{} {
	ret, _ := c.baseClient.RPush(key, value).Result()
	return ret
}

func (c Client) LPop(key string) interface{} {
	ret, _ := c.baseClient.LPop(key).Result()
	return ret
}

func (c Client) LLen(key string) interface{} {
	ret, _ := c.baseClient.LLen(key).Result()
	return ret
}

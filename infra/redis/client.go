package redis

import (
	"gimServer/conf"
	"gimServer/infra/utils"
	"github.com/go-redis/redis/v7"
)

func InitClient(config *conf.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})
	if e := client.Ping().Err(); e != nil {
		utils.Must(e)
	}
	return client
}

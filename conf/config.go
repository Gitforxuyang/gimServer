package conf

import (
	"fmt"
	"gimServer/infra/utils"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

type RabbitConfig struct {
}
type MongoConfig struct {
	Url         string
	MaxPoolSize uint64
	MinPoolSize uint64
}

type Config struct {
	Redis    *RedisConfig
	Rabbit   *RabbitConfig
	Mongo    *MongoConfig
	LogLevel string
}

func InitConfig() *Config {
	config := Config{}
	v := viper.New()
	v.SetConfigName("config.default")
	v.AddConfigPath("./conf")
	v.SetConfigType("json")

	err := v.ReadInConfig()
	utils.Must(err)
	v.BindEnv("ENV")
	env := v.GetString("ENV")
	if env == "" {
		env = "default"
	}
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	err = v.MergeInConfig()
	utils.Must(err)
	err = v.UnmarshalKey("redis", &config.Redis)
	utils.Must(err)
	err = v.UnmarshalKey("rabbit", &config.Rabbit)
	utils.Must(err)
	err = v.UnmarshalKey("mongo", &config.Mongo)
	utils.Must(err)
	config.LogLevel = v.GetString("logLevel")
	return &config
}

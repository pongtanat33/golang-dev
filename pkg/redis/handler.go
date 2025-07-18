package redis

import (
	"bufferbox_backend_go/logs"
	"context"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	Rdb *redis.ClusterClient
	Ctx = context.Background()
)

func InitRedis() error {
	env := os.Args[1]
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	logs.Info("Initializing Redis connection...")
	logs.Info("Redis Host: " + viper.GetString("REDIS.HOST"))
	logs.Info("Redis User: " + viper.GetString("REDIS.USER"))
	logs.Info("Redis Password: " + viper.GetString("REDIS.PASSWORD"))

	Rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{viper.GetString("REDIS.HOST")}, // สามารถใส่หลาย node ได้
		Username: viper.GetString("REDIS.USER"),
		Password: viper.GetString("REDIS.PASSWORD"),
	})

	// // //ตรวจสอบการเชื่อมต่อ
	if err := Rdb.Ping(Ctx).Err(); err != nil {
		return err
	}
	return nil
}

package drivers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func RunRedisDB() error {
	client := redis.NewClient(&redis.Options{
		Addr:     Viper.GetString("ResisDB.IP") + ":" + Viper.GetString("ResisDB.Port"),
		Password: "",                         // no password set
		DB:       Viper.GetInt("ResisDB.DB"), // use default DB
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	fmt.Println("RedisDB OK:" + pong)

	RedisDB = client
	return nil
}

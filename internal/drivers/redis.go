package drivers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func RunRedisDB() {
	client := redis.NewClient(&redis.Options{
		Addr:     Viper.GetString("ResisDB.IP") + ":" + Viper.GetString("ResisDB.Port"),
		Password: "",                         // no password set
		DB:       Viper.GetInt("ResisDB.DB"), // use default DB
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic("RedisDB error:" + err.Error())
	}

	fmt.Println("RedisDB OK:" + pong)

	RedisDB = client
}

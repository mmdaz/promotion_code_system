package redis

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
)

type Redis struct {
	Client *redis.Client
}

type Option struct {
	Host       string
	Port       string
	PoolSize   int
	DB         int
	Pass       string
	MaxRetries int
}

func NewRedisWithOption(enable bool, option Option) *Redis {
	var redisClient *redis.Client
	if enable {
		redisClient = redis.NewClient(&redis.Options{
			Addr:       option.Host + ":" + option.Port,
			MaxRetries: option.MaxRetries,
			PoolSize:   option.PoolSize,
			Password:   option.Pass,
			DB:         option.DB,
		})
	} else {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr:     s.Addr(),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Failed to create redis client")
	}

	log.Debug("Pong is here:", pong)

	return &Redis{
		Client: redisClient,
	}
}


package redis

import (
	"github.com/go-redis/redis/v7"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v7"
	"github.com/labstack/gommon/log"
)

type Redis struct {
	Client   *redis.Client
	SyncPool redsyncredis.Pool
}

type Option struct {
	Host       string
	Port       string
	PoolSize   int
	DB         int
	Pass       string
	MaxRetries int
}

func NewRedisWithOption(option Option) *Redis {
	var redisClient *redis.Client
	redisClient = redis.NewClient(&redis.Options{
		Addr:       option.Host + ":" + option.Port,
		MaxRetries: option.MaxRetries,
		PoolSize:   option.PoolSize,
		Password:   option.Pass,
		DB:         option.DB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Failed to create redis client")
	}
	log.Debug("Pong is here:", pong)

	pool := goredis.NewPool(redisClient)

	return &Redis{
		Client:   redisClient,
		SyncPool: pool,
	}
}

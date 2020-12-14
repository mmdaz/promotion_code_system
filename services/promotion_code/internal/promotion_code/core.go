package promotion_code

import (
	"gitlab.com/mmdaz/arvan-challenge/pkg/redis"
)

type Core struct {
	redis *redis.Redis
}

func NewCore(redis *redis.Redis) *Core {
	return &Core{redis: redis}
}

func (c *Core) ApplyPromotionCode(phoneNumber string) error {

	return nil
}

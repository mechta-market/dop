package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mechta-market/dop/adapters/logger"
)

type St struct {
	lg     logger.WarnAndError
	prefix string

	r   *redis.Client
	ctx context.Context
}

func New(lg logger.WarnAndError, url, psw string, db int, prefix string) *St {
	return &St{
		lg:     lg,
		prefix: prefix,

		r: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: psw,
			DB:       db,
		}),
		ctx: context.Background(),
	}
}

func (c *St) Get(key string) ([]byte, bool, error) {
	data, err := c.r.Get(c.ctx, c.prefix+key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		c.lg.Errorw("Redis: fail to 'get'", err)
		return nil, false, err
	}

	return data, true, nil
}

func (c *St) GetJsonObj(key string, dst any) (bool, error) {
	dataRaw, ok, err := c.Get(key)
	if err != nil || !ok {
		return ok, err
	}

	err = json.Unmarshal(dataRaw, dst)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *St) Set(key string, value []byte, expiration time.Duration) error {
	err := c.r.Set(c.ctx, c.prefix+key, value, expiration).Err()
	if err != nil {
		c.lg.Errorw("Redis: fail to 'set'", err)
	}

	return err
}

func (c *St) SetJsonObj(key string, value any, expiration time.Duration) error {
	dataRaw, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Set(key, dataRaw, expiration)
}

func (c *St) Del(key string) error {
	err := c.r.Del(c.ctx, c.prefix+key).Err()
	if err != nil {
		c.lg.Errorw("Redis: fail to 'del'", err)
	}

	return err
}

func (c *St) Keys(pattern string) []string {
	var err error
	var cursor uint64
	var keys []string

	resKeys := make([]string, 0)
	for {
		keys, cursor, err = c.r.Scan(c.ctx, cursor, c.prefix+pattern, 30).Result()
		if err != nil {
			break
		}
		resKeys = append(resKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	return resKeys
}

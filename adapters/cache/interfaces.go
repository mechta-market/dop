package cache

import "time"

type Cache interface {
	Get(key string) ([]byte, bool, error)
	GetJsonObj(key string, dst any) (bool, error)
	Set(key string, value []byte, expiration time.Duration) error
	SetJsonObj(key string, value any, expiration time.Duration) error
	Del(key string) error
	Keys(pattern string) []string
}

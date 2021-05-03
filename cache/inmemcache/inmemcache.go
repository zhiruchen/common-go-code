package inmemcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	// create kvstore with default expiration: 5 mins, and clean up interval: 10 mins
	kvstore = cache.New(5*time.Minute, 10*time.Minute)
)

func Set(key string, value interface{}, d time.Duration) {
	kvstore.Set(key, value, d)
}

func Get(key string) (value interface{}, ok bool) {
	return kvstore.Get(key)
}

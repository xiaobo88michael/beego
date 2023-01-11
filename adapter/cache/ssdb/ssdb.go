package ssdb

import (
	"github.com/xiaobo88michael/beego/adapter/cache"
	ssdb2 "github.com/xiaobo88michael/beego/client/cache/ssdb"
)

// NewSsdbCache create new ssdb adapter.
func NewSsdbCache() cache.Cache {
	return cache.CreateNewToOldCacheAdapter(ssdb2.NewSsdbCache())
}

func init() {
	cache.Register("ssdb", NewSsdbCache)
}

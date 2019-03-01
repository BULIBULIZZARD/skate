package cache

import (
	"file/skate/config"
	"github.com/muesli/cache2go"
	"time"
)

type MyCache struct {

}

func NewCache() *MyCache {
	return new(MyCache)
}

func (c *MyCache) getCache2go() *cache2go.CacheTable {
	cache := cache2go.Cache(config.GetConfig().GetCacheTableName())
	//cache.SetAddedItemCallback(func(entry *cache2go.CacheItem) {
	//	//log.Print("Added:  ", entry.Key(), "   ",entry.CreatedOn())
	//})
	return cache
}
func (c *MyCache) SetCache(key string, value string) {
	cache := c.getCache2go()
	cache.Add(key,3600*time.Second,value)
}
func (c *MyCache) GetCache(key string) interface{} {
	cache :=c.getCache2go()
	res, err := cache.Value(key)
	if err!=nil{
		return ""
	}
	return res.Data()
}

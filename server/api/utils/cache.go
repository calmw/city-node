package utils

import (
	"sync"
	"time"
)

var EventCache *Cache

type Cache struct {
	Lock *sync.RWMutex
	Key  []string
	Data map[string][]byte
	Time map[string]map[string]int64
}

func InitCache() {
	EventCache = &Cache{
		Lock: &sync.RWMutex{},
		Key:  []string{},
		Data: make(map[string][]byte),
		Time: make(map[string]map[string]int64),
	}

	t := time.Tick(time.Minute * 10)
	for {
		select {
		case <-t:
			EventCache.CheckExpired()
		}
	}
}

func (c *Cache) Set(key string, data []byte, duration int64) {
	ok, _ := c.Get(key)
	if !ok {
		c.Key = append(c.Key, key)
	}
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Data[key] = data
	if c.Time[key] == nil {
		c.Time[key] = map[string]int64{}
	}
	c.Time[key]["ctime"] = time.Now().Unix()
	c.Time[key]["duration"] = duration
}

func (c *Cache) Get(key string) (bool, []byte) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	data, ok := c.Data[key]
	return ok, data
}

func (c *Cache) CheckExpired() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	newKey := c.Key
	for _, k := range c.Key {
		ctime := c.Time[k]["ctime"]
		duration := c.Time[k]["duration"]
		if time.Now().Unix()-ctime > duration {
			// 删除key
			for j, nk := range newKey {
				if nk == k {
					newKey = append(newKey[:j], newKey[j+1:]...)
				}
			}
			// 删除 data
			_, hasData := c.Data[k]
			if hasData {
				delete(c.Data, k)
			}
			//删除 time
			_, hasTime := c.Time[k]
			if hasTime {
				delete(c.Time, k)
			}
		}
	}
	c.Key = newKey

}

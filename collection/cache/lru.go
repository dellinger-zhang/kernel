package cache

import (
	"container/list"
	"time"
)

// LRUCache
type LRUCache struct {
	baseCache
	items     map[interface{}]*list.Element
	evictList *list.List
}

func newLRUCache(cb *CacheBuilder) *LRUCache {
	c := &LRUCache{}
	buildCache(&c.baseCache, cb)

	c.init()
	c.loadGroup.cache = c
	return c
}

func (c *LRUCache) init() {
	c.evictList = list.New()
	c.items = make(map[interface{}]*list.Element, c.size+1)
}

func (c *LRUCache) set(key, value interface{}) (interface{}, error) {
	var err error
	if c.serializeFunc != nil {
		value, err = c.serializeFunc(key, value)
		if err != nil {
			return nil, err
		}
	}

	var item *lruItem
	if it, ok := c.items[key]; ok {
		c.evictList.MoveToFront(it)
		item = it.Value.(*lruItem)
		item.value = value
	} else {
		if c.evictList.Len() >= c.size {
			c.evict(1)
		}
		item = &lruItem{
			clock: c.clock,
			key:   key,
			value: value,
		}
		c.items[key] = c.evictList.PushFront(item)
	}

	if c.expiration != nil {
		t := c.clock.Now().Add(*c.expiration)
		item.expiration = &t
	}

	if c.addedFunc != nil {
		c.addedFunc(key, value)
	}

	return item, nil
}

// set a new key-value pair
func (c *LRUCache) Set(key, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.set(key, value)
	return err
}

// Set a new key-value pair with an expiration time
func (c *LRUCache) SetWithExpire(key, value interface{}, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, err := c.set(key, value)
	if err != nil {
		return err
	}

	t := c.clock.Now().Add(expiration)
	item.(*lruItem).expiration = &t
	return nil
}

// Get a value from cache pool using key if it exists.
func (c *LRUCache) Get(key interface{}) (interface{}, error) {
	v, err := c.get(key, false)
	if err == KeyNotFoundError {
		return c.getWithLoader(key, true)
	}
	return v, err
}

// Get a value from cache pool using key if it exists.
func (c *LRUCache) GetIFPresent(key interface{}) (interface{}, error) {
	v, err := c.get(key, false)
	if err == KeyNotFoundError {
		return c.getWithLoader(key, false)
	}
	return v, err
}

func (c *LRUCache) get(key interface{}, onLoad bool) (interface{}, error) {
	v, err := c.getValue(key, onLoad)
	if err != nil {
		return nil, err
	}
	if c.deserializeFunc != nil {
		return c.deserializeFunc(key, v)
	}
	return v, nil
}

func (c *LRUCache) getValue(key interface{}, onLoad bool) (interface{}, error) {
	c.mu.Lock()
	item, ok := c.items[key]
	if ok {
		it := item.Value.(*lruItem)
		if !it.IsExpired(nil) {
			c.evictList.MoveToFront(item)
			v := it.value
			c.mu.Unlock()
			if !onLoad {
				c.stats.IncrementHitCount()
			}
			return v, nil
		}
		c.removeElement(item)
	}
	c.mu.Unlock()
	if !onLoad {
		c.stats.IncrementMissCount()
	}
	return nil, KeyNotFoundError
}

func (c *LRUCache) getWithLoader(key interface{}, isWait bool) (interface{}, error) {
	if c.loaderExpireFunc == nil {
		return nil, KeyNotFoundError
	}
	value, _, err := c.load(key, func(v interface{}, expiration *time.Duration, e error) (interface{}, error) {
		if e != nil {
			return nil, e
		}
		c.mu.Lock()
		defer c.mu.Unlock()
		item, err := c.set(key, v)
		if err != nil {
			return nil, err
		}
		if expiration != nil {
			t := c.clock.Now().Add(*expiration)
			item.(*lruItem).expiration = &t
		}
		return v, nil
	}, isWait)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// evict removes the oldest item from the cache.
func (c *LRUCache) evict(count int) {
	for i := 0; i < count; i++ {
		ent := c.evictList.Back()
		if ent == nil {
			return
		} else {
			c.removeElement(ent)
		}
	}
}

// Removes the provided key from the cache.
func (c *LRUCache) Remove(key interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.remove(key)
}

func (c *LRUCache) remove(key interface{}) bool {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

func (c *LRUCache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	entry := e.Value.(*lruItem)
	delete(c.items, entry.key)
	if c.evictedFunc != nil {
		entry := e.Value.(*lruItem)
		c.evictedFunc(entry.key, entry.value)
	}
}

func (c *LRUCache) keys() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]interface{}, len(c.items))
	var i = 0
	for k := range c.items {
		keys[i] = k
		i++
	}
	return keys
}

// Returns a slice of the keys in the cache.
func (c *LRUCache) Keys() []interface{} {
	keys := []interface{}{}
	for _, k := range c.keys() {
		_, err := c.GetIFPresent(k)
		if err == nil {
			keys = append(keys, k)
		}
	}
	return keys
}

// Returns all key-value pairs in the cache.
func (c *LRUCache) GetALL() map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	for _, k := range c.keys() {
		v, err := c.GetIFPresent(k)
		if err == nil {
			m[k] = v
		}
	}
	return m
}

// Returns the number of items in the cache.
func (c *LRUCache) Len() int {
	return len(c.GetALL())
}

// Completely clear the cache
func (c *LRUCache) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.init()
}

type lruItem struct {
	clock      Clock
	key        interface{}
	value      interface{}
	expiration *time.Time
}

// returns boolean value whether this item is expired or not.
func (it *lruItem) IsExpired(now *time.Time) bool {
	if it.expiration == nil {
		return false
	}
	if now == nil {
		t := it.clock.Now()
		now = &t
	}
	return it.expiration.Before(*now)
}
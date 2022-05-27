package cache

import "time"

type Item struct {
	value    string
	deadline time.Time
}

type Cache struct {
	items map[string]Item
}

func NewCache() Cache {
	return Cache{items: make(map[string]Item)}
}

func (c *Cache) Get(key string) (string, bool) {
	if item, ok := c.items[key]; ok {
		if item.deadline.IsZero() || item.deadline.After(time.Now()) {
			return item.value, true
		}
		return "", false
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.items[key] = Item{value, time.Time{}}
}

func (c *Cache) Keys() []string {
	var keys []string
	for k, v := range c.items {
		if v.deadline.IsZero() || v.deadline.After(time.Now()) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = Item{value, deadline}
}

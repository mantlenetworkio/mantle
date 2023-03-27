package types

import "github.com/elliotchance/orderedmap/v2"

type Cache[K comparable, V any] struct {
	size  int
	cache *orderedmap.OrderedMap[K, V]
}

func NewCache[K comparable, V any](size int) *Cache[K, V] {
	cache := orderedmap.NewOrderedMap[K, V]()
	return &Cache[K, V]{
		size:  size,
		cache: cache,
	}
}

func (c *Cache[K, V]) Set(key K, value V) (didSet bool) {
	if c.cache.Len() == c.size {
		el := c.cache.Front()
		bol := c.cache.Delete(el.Key)
		if !bol {
			return bol
		}
	}
	return c.cache.Set(key, value)
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	return c.cache.Get(key)
}

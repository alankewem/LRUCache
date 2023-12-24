package lru

import (
	"container/list"
)

type LRUCache struct {
	capacity uint
	items    map[string]*list.Element
	order    *list.List
}

func New(capacity uint) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (l *LRUCache) Get(key string) (value interface{}, exists bool) {
	if item, exists := l.items[key]; exists {
		l.order.MoveToFront(item)
		return item.Value, exists
	}

	return nil, false
}

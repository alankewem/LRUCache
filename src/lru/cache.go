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

func (lru *LRUCache) Add(key string, value interface{}) {
	if uint(len(lru.items)) >= lru.capacity {
		back := lru.order.Back()
		if back != nil {
			delete(lru.items, back.Value.(map[string]interface{})["key"].(string))
			lru.order.Remove(back)
		}
	}

	item := map[string]interface{}{"key": key, "value": value}
	lru.items[key] = lru.order.PushFront(item)
}

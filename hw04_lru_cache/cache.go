package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
// 	key   Key
// 	value interface{}
// }

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, inMap := l.items[key]; inMap {
        // элемент присутствовал в кеше
		item.Value = value
		// fmt.Println("Key were in cache", l.queue)
		l.queue.MoveToFront(item)
        return true
	} else {
		// fmt.Println("Added a new key", l.queue)
		if l.queue.Len() == l.capacity {
            fmt.Println("Элемент превышает длину")
			// в этом случае надо удалить последний элемент из очереди и его значение из словаря
			oldLast := l.queue.Back()
			l.queue.Remove(oldLast)
			nLast := new(ListItem)
			*oldLast = *nLast
		} 
		// Просто добавить
        newFirst := l.queue.PushFront(value)
        l.items[key] = newFirst
		return false
	}
}


func (l *lruCache) Clear() {
    l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}


func (l *lruCache) Get(key Key) (interface{}, bool) {
	if v, inMap := l.items[key]; inMap {
        fmt.Println("Есть элемент", v)
		if l.queue.Front() == v {
			return v.Value, inMap
		}
		nilItem := new(ListItem)
		if *v != *nilItem {
			l.queue.MoveToFront(v)
		} else {
			delete(l.items, key)
		}
		return v.Value, inMap
	} else {
        fmt.Println("нет элемента")
		return nil, false
	}
}


func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

package hw04lrucache

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

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, inMap := l.items[key]; inMap {
		item.Value = value
		l.queue.MoveToFront(item)
		return true
	}
	if l.queue.Len() == l.capacity {
		// в этом случае надо удалить последний элемент из очереди и его значение из словаря
		oldLast := l.queue.Back()
		l.queue.Remove(oldLast)
		// Поскольку у нас нет ключа последнего элемента в мапе, то заменяем значение новым ListItem
		nLast := new(ListItem)
		*oldLast = *nLast
	}
	newFirst := l.queue.PushFront(value)
	l.items[key] = newFirst
	return false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if v, inMap := l.items[key]; inMap {
		if l.queue.Front() == v {
			return v.Value, inMap
		}
		nilItem := new(ListItem)
		if *v != *nilItem {
			l.queue.MoveToFront(v)
		} else {
			// если по ключу мы получили пустую ListItem, то это значение было вытолкнуто ранее и можно удалять его из мапы
			delete(l.items, key)
		}
		return v.Value, inMap
	}
	return nil, false
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

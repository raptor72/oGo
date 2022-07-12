package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	first  *ListItem
	last   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.first
}

func (l *list) Back() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	oldFirst := l.Front()
	newFirst := ListItem{Value: v, Prev: nil, Next: oldFirst}
	l.first = &newFirst
	if oldFirst != nil {
		oldFirst.Prev = &newFirst
	} else {
		l.first = &newFirst
		l.last = &newFirst
	}
	l.length++
	return &newFirst
}

func (l *list) PushBack(v interface{}) *ListItem {
	oldLast := l.Back()
	newLast := ListItem{Value: v, Prev: oldLast, Next: nil}
	l.last = &newLast
	if oldLast != nil {
		oldLast.Next = &newLast
	} else {
		l.first = &newLast
		l.last = &newLast
	}
	l.length++
	return &newLast
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.first == i:
		l.first = l.first.Next
		l.first.Prev = nil
	case l.last == i:
		l.last = l.last.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	switch i {
	case l.first:
		return
	case l.last:
		l.last = l.Back().Prev
	default:
		i.Next.Prev = i.Prev
	}
	oldFront := l.Front()
	oldFront.Prev = i
	i.Prev.Next = i.Next
	i.Prev = nil
	i.Next = l.Front()
	l.first = i
}

func NewList() List {
	return new(list)
}

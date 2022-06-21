package hw04lrucache

// package main
// import (
// "fmt"
// )

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
	first *ListItem
	last *ListItem
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
    // нужно сделать тест на использование этого метода на пустом списке и на не пустом списке
	oldFirst := l.Front()
	newFirst := ListItem{Value: v, Prev: nil, Next: oldFirst}
	l.first = &newFirst
	if oldFirst != nil {
        oldFirst.Prev = &newFirst
	} else {
		l.first = &newFirst
		l.last = &newFirst
	}
	l.length ++
    return &newFirst
}

func (l *list) PushBack(v interface{}) *ListItem {
    // нужно сделать тест на использование этого метода на пустом списке и на не пустом списке
	oldLast := l.Back()
	newLast := ListItem{Value: v, Prev: oldLast, Next: nil}
    l.last = &newLast
    if oldLast != nil {
        oldLast.Next = &newLast
	} else {
        l.first = &newLast
        l.last = &newLast
	}
	l.length ++
	return &newLast
}

func (l *list) Remove(i *ListItem) {
    if l.first == i {
        l.first = l.first.Next
        l.first.Prev = nil
	} else if l.last == i {
        l.last = l.last.Prev
        l.last.Next = nil
	} else {
		// Общий случай
		i.Prev.Next = i.Next
        i.Next.Prev = i.Prev
	}
    l.length--
}


func (l *list) MoveToFront(i *ListItem) {
    if i == l.first {
		return
	} else if i == l.last {
		l.last = l.Back().Prev
	} else { 
		i.Next.Prev = i.Prev
	}
	// Это общая для всех часть
	old_front := l.Front()
    old_front.Prev = i
	i.Prev.Next = i.Next // тут будет nil если i в конце списка и так и надо
    i.Prev = nil
    i.Next = l.Front()
    l.first = i

}

func NewList() List {
	return new(list)
}

// func main() {
//     l := NewList()
//      fmt.Println(l)
// 	fmt.Println(l.Len())   
//     l.PushFront(10) // [10]
//     l.PushFront(5)  // [5, 10]
//     l.PushFront(2)  // [2, 5, 10]
// 	fmt.Println(l)
//     fmt.Println(l.Front().Value)
//     fmt.Println(l.Back().Value)
// 	fmt.Println(l.Len())
//     l.MoveToFront(l.Back()) // [10, 2, 5]
//     l.MoveToFront(l.Back()) // [5, 10, 2]
//     l.MoveToFront(l.Back())
// 	fmt.Println(l)                   // &{3 0xc0000c0040 0xc0000c0000}
//     fmt.Println(l.Front())           // &{2 0xc00006a040 <nil>}
//     fmt.Println(l.Front().Next)      // &{5 0xc00006a020 0xc00006a060}
//     fmt.Println(l.Front().Next.Next) // &{10 <nil> 0xc00006a040}
//     fmt.Println(l.Front().Prev)      // <nil>
//     fmt.Println(l.Back().Next)       // <nil>
// }
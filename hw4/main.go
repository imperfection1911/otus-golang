package main

import (
	"fmt"
)

// структура item
type Item struct {
	value interface{}
	next, prev  *Item
}

func (i Item) Value() interface{} {
	return i.value
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}

type List struct {
	first *Item
	last *Item
	len int
}

func (l List) Len() int {
	return l.len
}

func (l List) First() *Item {
	return l.first
}

func (l List) Last() *Item {
	return l.last
}

func (l *List) PushBack(value interface{})  {
	item := &Item{value: value}
	if l.last == nil {
		l.last = item
		if l.first != nil {
			l.first.next = item
		} else {
			l.first = item
		}
	} else {
		l.last.next = item
		item.prev = l.last
	}
	l.len ++
	l.last = item
}

func (l *List) PushFront(value interface{}) {
	item := &Item{value: value}
	if l.first == nil {
		l.first = item
		if l.last != nil {
			l.last.prev = item
		} else {
			l.last = item
		}
	} else {
		l.first.prev = item
		item.next = l.first
	}
	l.len ++
	l.first = item
}

func (l *List) Remove(item *Item)  {
	if item.prev != nil {
		previous := item.prev
		previous.next = item.next
	}
	if item.next != nil {
		next := item.next
		next.prev = item.prev
	}
	if l.first == item {
		l.first = item.next
	}
	if l.last == item {
		l.last = item.prev
	}
	l.len--
}


func main()  {
	l := new(List)
	l.PushBack("atlanta")
	l.PushFront("Moscow")
	l.PushBack("LA")
	fmt.Printf("%+v\n", l)
	fmt.Printf("%+v\n", l.first)
	fmt.Printf("%+v\n", l.last.prev)
	fmt.Printf("%+v\n", l.last)
	l.Remove(l.last)
	fmt.Println("deleted")
	fmt.Printf("%+v\n", l)
	fmt.Printf("%+v\n", l.first)
	fmt.Printf("%+v\n", l.last)
	//fmt.Printf("%+v\n", l)
}
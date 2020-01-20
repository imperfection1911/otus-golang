package main

import (
	"fmt"
)

//Item of the double linked list
type Item struct {
	value      interface{}
	next, prev *Item
}

// Value returns Item value attribute.
func (i Item) Value() interface{} {
	return i.value
}

//Next returns Item next attribute
func (i Item) Next() *Item {
	return i.next
}

//Prev returns Item Prev attribute
func (i Item) Prev() *Item {
	return i.prev
}

// Не понимаю необходимости реализации 3 метордов выше, если всегда могу получить value как item.value например. :(

//List struct of dll
type List struct {
	first *Item
	last  *Item
	len   int
}

//Len returns list len
func (l List) Len() int {
	return l.len
}

//First returns first item in list
func (l List) First() *Item {
	return l.first
}

//Last returns last item in list
func (l List) Last() *Item {
	return l.last
}

//PushBack push Item in the end of list
func (l *List) PushBack(value interface{}) {
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
	l.len++
	l.last = item
}

//PushFront push item in fron of list
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
	l.len++
	l.first = item
}

//Remove suddenly removes item friom list
func (l *List) Remove(item *Item) {
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

func main() {
	l := new(List)
	l.PushBack("A")
	l.PushFront("B")
	l.PushBack("C")
	fmt.Printf("%#v\n", l.last)
	// вот тоже не до конца понял. Remove ждет указатель, но если передать l.last как &l.last не компилируется.
	// Понимаю что l.last уже указатель, но забыл как проверить тип переменной:( Легко запутаться
	l.Remove(l.last)
	fmt.Printf("%#v\n", l.last)
	//fmt.Printf("%+v\n", l)
}

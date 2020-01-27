package main

import (
	"errors"
	"fmt"
)

//Item of the double linked list
type Item struct {
	value      interface{}
	next, prev *Item
	list     *List
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
	item := &Item{value: value, list: l}
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
	item := &Item{value: value, list: l}
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
func (l *List) Remove(item *Item) (error) {
	var err error
	if item.list == l {
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
		item.list = nil
	} else {
		err = errors.New("item is from another list or have been already deleted")
	}
	return err
}

func main() {
	l := new(List)
	l.PushBack("A")
	l.PushFront("B")
	l.PushBack("C")
	fmt.Printf("%#v\n", l.last)
	err := l.Remove(l.last)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", l.last)
}

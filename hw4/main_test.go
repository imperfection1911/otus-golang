package main

import (
	"testing"
)

func TestItem_Value(t *testing.T) {
	value := "test"
	item := &Item{value: value}
	if item.Value() != value {
		t.Fatalf("test failed %#v not equal %s", item, value)
	}
}

func TestItem_Next(t *testing.T) {
	item1 := &Item{value: "A"}
	item2 := &Item{value: "B"}
	item1.next = item2
	next := item1.Next()
	if next.value != "B" {
		t.Fatalf("test failed %#v", next)
	}
}

func TestItem_Prev(t *testing.T) {
	item1 := &Item{value: "A"}
	item2 := &Item{value: "B"}
	item1.prev = item2
	prev := item1.Prev()
	if prev.value != "B" {
		t.Fatalf("test failed %#v", prev)
	}
}

func TestList_Len(t *testing.T) {
	l := &List{}
	l.GenerateID()
	l.PushFront("A")
	l.PushFront("B")
	l.PushFront("C")
	len := l.Len()
	if len != 3 {
		t.Fatalf("test failed %v not equal 3", len)
	}

}

// На First и Last тесты писать не стал, они очевидны

func TestList_PushBack(t *testing.T) {
	l := &List{}
	l.GenerateID()
	l.PushFront("A")
	l.PushFront("B")
	l.PushFront("C")
	l.PushBack("D")
	if l.last.Value() != "D" {
		t.Fatalf("%#v", l.last)
	}
}

func TestList_PushFront(t *testing.T) {
	l := &List{}
	l.GenerateID()
	l.PushBack("A")
	l.PushBack("B")
	l.PushBack("C")
	l.PushFront("D")
	if l.first.Value() != "D" {
		t.Fatalf("%#v", l.first)
	}
}

func TestList_Remove(t *testing.T) {
	l := &List{}
	l.GenerateID()
	l.PushBack("A")
	l.PushBack("B")
	l.PushBack("C")
	l.PushFront("D")
	_, err := l.Remove(l.first)
	if err != nil {
		t.Fatalf("%#v", err)
	}
	if l.first.value != "A" {
		t.Fatalf("%#v", l.first)
	}
}

func TestList_Remove_deleted(t *testing.T) {
	l := &List{}
	l.GenerateID()
	i := &Item{value: "A", deleted: true}
	_, err := l.Remove(i)
	if err == nil {
		t.Fatalf("%#v", err)
	}
}

func TestList_Remove_another(t *testing.T) {
	l := &List{}
	l.GenerateID()
	i := &Item{value: "A", listID: "1qaz2wsx"}
	_, err := l.Remove(i)
	if err == nil {
		t.Fatalf("%#v", err)
	}
}

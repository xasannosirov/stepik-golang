package main

import (
	"container/list"
	// "fmt"
)

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()

	for e := l.Back(); e != nil; e = e.Prev() {
		reversedList.PushBack(e.Value)
	}

	return reversedList
}
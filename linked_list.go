package main

// List represents a singly-linked list that holds
// values of any type.

import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) hasNext() bool {
	if nil == list {
		return false
	}
	if nil == list.next {
		return false
	}
	return true
}

func (list *List[T]) iterator() func() *List[T] {
	iter := list
	return func() *List[T] {
		iter = iter.next
		return iter
	}
}

func main() {
	var list *List[int] = &List[int]{
		val: 0,
		next: &List[int]{
			val: 2,
			next: &List[int]{
				val: 4,
				next: &List[int]{
					val:  5,
					next: nil,
				},
			},
		},
	}

	iter := list.iterator()
	for list.hasNext() {
		fmt.Println(list.val)
		list = iter()
	}
}

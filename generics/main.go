package main

import (
	"fmt"
	"strings"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	if l == nil {
		return "[]"
	}
	var b strings.Builder
	b.WriteString("[")
	for cur := l; cur != nil; cur = cur.next {
		fmt.Fprintf(&b, "%v", cur.val)
		if cur.next != nil {
			b.WriteString(" -> ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List[T]) last() *List[T] {
	if l == nil {
		return nil
	}
	cur := l
	for cur.next != nil {
		cur = cur.next
	}

	return cur
}

func (l *List[T]) append(list *List[T]) *List[T] {
	if l == nil {
		return list
	}

	last := l.last()
	last.next = list
	return l
}

type Point struct {
	x, y int
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	var l *List[int]
	l = l.append(&List[int]{val: 1})
	l = l.append(&List[int]{val: 2})

	fmt.Println(l)
	fmt.Println(l.last())
}

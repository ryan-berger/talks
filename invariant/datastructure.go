

package main

// START INTRO OMIT
type node[T any] {
	next *node[T]
	val T
}

type linked[T any] struct {
	head *node[T]
}

func (l *linked[T]) insert(n *node[T]) {
	if l.head == nil {
		l.head = n
		return
	}
	var prev, cur *node[T]
	cur = l.head
	for cur != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = n
}
// END INTRO OMIT

// START BETTER OMIT
func (l *linked[T]) insertBetter(n *node[T]) {
	insertionPoint := &l.head
	for *insertionPoint != nil {
		insertionPoint = &(*insertionPoint).next
	}
	*insertionPoint = n
}
// END BETTER OMIT



func main() {}




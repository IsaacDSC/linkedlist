package algoritms

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Unshift(value T) {
	node := &Node[T]{value, l.head}
	l.head = node
}

func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{value, nil}
	if l.head == nil {
		l.head = node
		return
	}
	head := l.head
	for head.next != nil {
		head = head.next
	}
	head.next = node
}

func (l *LinkedList[T]) AppendOnIndex(value T, index int) {
	if index == 0 {
		l.Unshift(value)
		return
	}

	node := &Node[T]{value, nil}
	actual := l.head
	for i := 0; actual != nil && i < index-1; i++ {
		actual = actual.next
	}

	if actual == nil {
		fmt.Println("Posição inválida")
		return
	}

	node.next = actual.next
	actual.next = node
}

func (l *LinkedList[T]) Search(value T) (int, bool) {
	head := l.head
	index := 0
	for head != nil {
		if head.value == value {
			return index, true
		}
		head = head.next
		index++
	}
	return -1, false
}

func (l *LinkedList[T]) Delete(value T) bool {
	current := l.head
	var anterior *Node[T]

	for current != nil {
		if current.value == value {
			if anterior == nil {
				l.head = current.next
			} else {
				anterior.next = current.next
			}
			return true
		}
		anterior = current
		current = current.next
	}
	return false
}

func (l *LinkedList[T]) Debug() {
	head := l.head
	for head != nil {
		fmt.Printf("%v -> ", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

package algoritms

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Unshift(value int) {
	node := &Node{value, l.head}
	l.head = node
}

func (l *LinkedList) Append(value int) {
	node := &Node{value, nil}
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

func (l *LinkedList) AppendOnIndex(value int, index int) {
	if index == 0 {
		l.Unshift(value)
		return
	}

	node := &Node{value, nil}
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

func (l *LinkedList) Search(value int) (int, bool) {
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

func (l *LinkedList) Delete(value int) bool {
	current := l.head
	var anterior *Node

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

func (l *LinkedList) Debug() {
	head := l.head
	for head != nil {
		fmt.Printf("%d -> ", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

package algoritms

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T] // permite acessar o início da lista em O(1)
	tail *Node[T] // permite acessar o final da lista em O(1)
	size int      // rastreia o número de elementos para acesso O(1) ao tamanho
}

func (l *LinkedList[T]) Unshift(value T) {
	node := &Node[T]{value, l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{value, nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *LinkedList[T]) AppendOnIndex(value T, index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("invalid index: %d", index)
	}

	if index == 0 {
		l.Unshift(value)
		return
	}

	if index == l.size {
		l.Append(value)
		return
	}

	node := &Node[T]{value, nil}
	actual := l.head
	for i := 0; i < index-1; i++ {
		actual = actual.next
	}

	node.next = actual.next
	actual.next = node
	l.size++
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
				if l.head == nil {
					l.tail = nil // Lista ficou vazia
				}
			} else {
				anterior.next = current.next
				if current == l.tail {
					l.tail = anterior // Atualizando tail se removemos o último elemento
				}
			}
			l.size--
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
	fmt.Printf("Size: %d\n", l.size)
}

// O método Size() é uma interface pública para clientes da estrutura
// ele não deve ser utilizado internamente para um acesso mais eficiente use l.size
func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) FindByIndex(index int) (T, error) {
	var zero T

	// Validação de índice usando o campo size
	if index < 0 || index >= l.size {
		return zero, fmt.Errorf("índice %d fora dos limites (0-%d)", index, l.size-1)
	}

	// Otimização: verificar extremidades primeiro
	if index == 0 {
		return l.head.value, nil
	}

	if index == l.size-1 {
		return l.tail.value, nil
	}

	// Caso geral: percorrer a lista até o índice desejado
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

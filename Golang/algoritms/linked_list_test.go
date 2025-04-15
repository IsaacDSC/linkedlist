package algoritms

import (
	"testing"
)

// Testa a criação de uma lista vazia
func TestNewLinkedList(t *testing.T) {
	list := LinkedList[int]{}

	if list.head != nil {
		t.Errorf("Nova lista deveria ter head nil, mas tem %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("Nova lista deveria ter tail nil, mas tem %v", list.tail)
	}

	if list.size != 0 {
		t.Errorf("Nova lista deveria ter tamanho 0, mas tem %d", list.size)
	}
}

// Testa o método Append
func TestAppend(t *testing.T) {
	list := LinkedList[string]{}

	// Testa append em lista vazia
	list.Append("primeiro")
	if list.head.value != "primeiro" || list.tail.value != "primeiro" {
		t.Errorf("Falha ao adicionar primeiro elemento")
	}
	if list.size != 1 {
		t.Errorf("Tamanho deveria ser 1, mas é %d", list.size)
	}

	// Testa append em lista com elementos
	list.Append("segundo")
	list.Append("terceiro")

	if list.head.value != "primeiro" {
		t.Errorf("Head deveria ser 'primeiro', mas é %v", list.head.value)
	}

	if list.tail.value != "terceiro" {
		t.Errorf("Tail deveria ser 'terceiro', mas é %v", list.tail.value)
	}

	if list.size != 3 {
		t.Errorf("Tamanho deveria ser 3, mas é %d", list.size)
	}
}

// Testa o método Unshift
func TestUnshift(t *testing.T) {
	list := LinkedList[int]{}

	// Testa unshift em lista vazia
	list.Unshift(10)
	if list.head.value != 10 || list.tail.value != 10 {
		t.Errorf("Falha ao adicionar primeiro elemento com Unshift")
	}
	if list.size != 1 {
		t.Errorf("Tamanho deveria ser 1, mas é %d", list.size)
	}

	// Testa unshift em lista com elementos
	list.Unshift(20)
	list.Unshift(30)

	if list.head.value != 30 {
		t.Errorf("Head deveria ser 30, mas é %v", list.head.value)
	}

	if list.tail.value != 10 {
		t.Errorf("Tail deveria ser 10, mas é %v", list.tail.value)
	}

	if list.size != 3 {
		t.Errorf("Tamanho deveria ser 3, mas é %d", list.size)
	}
}

// Testa o método AppendOnIndex
func TestAppendOnIndex(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(10)
	list.Append(30)

	// Testa inserção no meio
	list.AppendOnIndex(20, 1)

	val, err := list.FindByIndex(1)
	if err != nil || val != 20 {
		t.Errorf("Esperado 20 na posição 1, obtido %v com erro %v", val, err)
	}

	// Testa inserção no início
	list.AppendOnIndex(5, 0)

	val, err = list.FindByIndex(0)
	if err != nil || val != 5 {
		t.Errorf("Esperado 5 na posição 0, obtido %v com erro %v", val, err)
	}

	// Testa inserção no final
	list.AppendOnIndex(40, list.size)

	val, err = list.FindByIndex(list.size - 1)
	if err != nil || val != 40 {
		t.Errorf("Esperado 40 na posição %d, obtido %v com erro %v", list.size-1, val, err)
	}

	// Testa índice inválido
	oldSize := list.size
	list.AppendOnIndex(100, -1) // Deve falhar silenciosamente
	if list.size != oldSize {
		t.Errorf("Tamanho não deveria mudar após índice inválido")
	}
}

// Testa o método Search
func TestSearch(t *testing.T) {
	list := LinkedList[string]{}
	list.Append("maçã")
	list.Append("banana")
	list.Append("laranja")

	// Testa buscar um elemento que existe
	index, found := list.Search("banana")
	if !found || index != 1 {
		t.Errorf("Esperado encontrar 'banana' no índice 1, obtido índice %d, encontrado: %v", index, found)
	}

	// Testa buscar um elemento que não existe
	index, found = list.Search("abacaxi")
	if found {
		t.Errorf("Não deveria encontrar 'abacaxi', mas encontrou no índice %d", index)
	}
}

// Testa o método Delete
func TestDelete(t *testing.T) {
	list := LinkedList[int]{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)

	// Deleta um elemento do meio
	success := list.Delete(20)
	if !success || list.size != 3 {
		t.Errorf("Falha ao deletar elemento do meio, success: %v, size: %d", success, list.size)
	}

	val, _ := list.FindByIndex(1)
	if val != 30 {
		t.Errorf("Após deletar 20, o índice 1 deveria ser 30, mas é %d", val)
	}

	// Deleta o primeiro elemento
	success = list.Delete(10)
	if !success || list.size != 2 {
		t.Errorf("Falha ao deletar primeiro elemento, success: %v, size: %d", success, list.size)
	}

	val, _ = list.FindByIndex(0)
	if val != 30 {
		t.Errorf("Após deletar 10, o índice 0 deveria ser 30, mas é %d", val)
	}

	// Deleta o último elemento
	success = list.Delete(40)
	if !success || list.size != 1 {
		t.Errorf("Falha ao deletar último elemento, success: %v, size: %d", success, list.size)
	}

	if list.tail.value != 30 {
		t.Errorf("Após deletar último, tail deveria ser 30, mas é %v", list.tail.value)
	}

	// Tenta deletar elemento que não existe
	success = list.Delete(50)
	if success {
		t.Errorf("Não deveria deletar elemento inexistente")
	}

	// Deleta o único elemento restante
	list.Delete(30)
	if list.head != nil || list.tail != nil || list.size != 0 {
		t.Errorf("Lista deveria estar vazia após deletar todos elementos")
	}
}

// Testa o método FindByIndex
func TestFindByIndex(t *testing.T) {
	list := LinkedList[string]{}
	list.Append("A")
	list.Append("B")
	list.Append("C")

	// Testa busca do primeiro elemento
	val, err := list.FindByIndex(0)
	if err != nil || val != "A" {
		t.Errorf("Esperado 'A' no índice 0, obtido %v com erro %v", val, err)
	}

	// Testa busca do último elemento
	val, err = list.FindByIndex(2)
	if err != nil || val != "C" {
		t.Errorf("Esperado 'C' no índice 2, obtido %v com erro %v", val, err)
	}

	// Testa busca de um elemento no meio
	val, err = list.FindByIndex(1)
	if err != nil || val != "B" {
		t.Errorf("Esperado 'B' no índice 1, obtido %v com erro %v", val, err)
	}

	// Testa busca com índice negativo
	_, err = list.FindByIndex(-1)
	if err == nil {
		t.Errorf("Deveria retornar erro para índice negativo")
	}

	// Testa busca com índice grande demais
	_, err = list.FindByIndex(3)
	if err == nil {
		t.Errorf("Deveria retornar erro para índice fora dos limites")
	}
}

// Teste de integração para verificar a estrutura completa
func TestIntegration(t *testing.T) {
	list := LinkedList[int]{}

	// Adiciona elementos de várias formas
	list.Append(30)
	list.Unshift(10)
	list.AppendOnIndex(20, 1)

	// Verifica se a ordem está correta
	expected := []int{10, 20, 30}
	for i, exp := range expected {
		val, err := list.FindByIndex(i)
		if err != nil || val != exp {
			t.Errorf("No índice %d, esperado %d, obtido %d com erro %v", i, exp, val, err)
		}
	}

	// Verifica o tamanho
	if list.Size() != 3 {
		t.Errorf("Tamanho esperado 3, obtido %d", list.Size())
	}

	// Verifica as extremidades
	if list.head.value != 10 || list.tail.value != 30 {
		t.Errorf("Extremidades incorretas, head: %v, tail: %v", list.head.value, list.tail.value)
	}

	// Remove um elemento e verifica a estrutura
	list.Delete(20)
	if list.Size() != 2 {
		t.Errorf("Após deleção, tamanho esperado 2, obtido %d", list.Size())
	}

	val, _ := list.FindByIndex(0)
	if val != 10 {
		t.Errorf("Após deleção, esperado 10 no índice 0, obtido %d", val)
	}

	val, _ = list.FindByIndex(1)
	if val != 30 {
		t.Errorf("Após deleção, esperado 30 no índice 1, obtido %d", val)
	}
}

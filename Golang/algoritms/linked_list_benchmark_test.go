package algoritms

import (
	"testing"
)

// ----- Benchmarks para operações de inserção no início -----

func BenchmarkLinkedList_Unshift(b *testing.B) {
	list := LinkedList[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Unshift(i)
	}
}

func BenchmarkSlice_Unshift(b *testing.B) {
	var slice []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Inserir no início de um slice requer deslocamento de todos elementos
		slice = append([]int{i}, slice...)
	}
}

// ----- Benchmarks para operações de inserção no final -----

func BenchmarkLinkedList_Append(b *testing.B) {
	list := LinkedList[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

func BenchmarkSlice_Append(b *testing.B) {
	var slice []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

// ----- Benchmarks para operações de inserção em posição específica -----

// Inserção no meio para diferentes tamanhos
func BenchmarkLinkedList_AppendOnIndex_Small(b *testing.B) {
	// Preparar uma lista com alguns elementos
	list := prepareLinkedList(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Inserir no meio
		list.AppendOnIndex(i, 50)
	}
}

func BenchmarkSlice_InsertMiddle_Small(b *testing.B) {
	// Preparar um slice com alguns elementos
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Inserir no meio
		middle := len(slice) / 2
		slice = append(slice[:middle], append([]int{i}, slice[middle:]...)...)
	}
}

func BenchmarkLinkedList_AppendOnIndex_Large(b *testing.B) {
	// Preparar uma lista maior
	list := prepareLinkedList(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Inserir no meio
		list.AppendOnIndex(i, 5000)
	}
}

func BenchmarkSlice_InsertMiddle_Large(b *testing.B) {
	// Preparar um slice maior
	slice := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Inserir no meio
		middle := len(slice) / 2
		slice = append(slice[:middle], append([]int{i}, slice[middle:]...)...)
	}
}

// ----- Benchmarks para operações de busca por valor -----

func BenchmarkLinkedList_Search(b *testing.B) {
	list := prepareLinkedList(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Buscar um valor que está próximo ao final (pior caso)
		list.Search(990)
	}
}

func BenchmarkSlice_Search(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Buscar um valor que está próximo ao final (pior caso)
		for j, v := range slice {
			if v == 990 {
				_ = j
				break
			}
		}
	}
}

// ----- Benchmarks para operações de busca por índice -----

func BenchmarkLinkedList_FindByIndex(b *testing.B) {
	list := prepareLinkedList(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Buscar um índice no meio da lista
		_, _ = list.FindByIndex(500)
	}
}

func BenchmarkSlice_GetByIndex(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Acesso direto por índice (O(1) para slices)
		_ = slice[500]
	}
}

// ----- Benchmarks para operações de remoção -----

func BenchmarkLinkedList_Delete_Beginning(b *testing.B) {
	// Prepara uma lista com 1000 elementos
	originalList := prepareLinkedList(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Restaura a lista ao estado inicial
		list := originalList
		list.head = originalList.head
		list.size = originalList.size

		// Remove o primeiro elemento
		list.Delete(list.head.value)
	}
}

func BenchmarkSlice_Delete_Beginning(b *testing.B) {
	// Prepara um slice com 1000 elementos
	originalSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		originalSlice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Restaura o slice ao estado inicial
		slice := originalSlice

		// Remove o primeiro elemento
		slice = slice[1:]
	}
}

func BenchmarkLinkedList_Delete_Middle(b *testing.B) {
	// Prepara uma lista com 1000 elementos
	originalList := prepareLinkedList(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Restaura a lista ao estado inicial
		list := originalList
		list.head = originalList.head
		list.size = originalList.size

		// Remover do meio
		list.Delete(500)
	}
}

func BenchmarkSlice_Delete_Middle(b *testing.B) {
	// Prepara um slice com 1000 elementos
	originalSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		originalSlice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Restaura o slice ao estado inicial
		slice := originalSlice

		// Remover do meio
		middle := 500
		slice = append(slice[:middle], slice[middle+1:]...)
	}
}

// ----- Funções auxiliares -----

// Prepara uma lista encadeada com n elementos
func prepareLinkedList(n int) LinkedList[int] {
	list := LinkedList[int]{}
	for i := 0; i < n; i++ {
		list.Append(i)
	}
	return list
}

// ----- Benchmark comparativo mostrando o impacto das otimizações -----

// Benchmark que demonstra o impacto da otimização do tail para operação Append
func BenchmarkLinkedList_AppendOptimized(b *testing.B) {
	// Usar a versão otimizada implementada
	list := LinkedList[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

// Simulação da implementação não-otimizada para comparação
func BenchmarkLinkedList_AppendNonOptimized(b *testing.B) {
	// Simular a versão não otimizada (sem tail)
	list := LinkedList[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simular o comportamento sem tail
		node := &Node[int]{i, nil}
		if list.head == nil {
			list.head = node
		} else {
			// Percorrer a lista para encontrar o último nó (O(n))
			current := list.head
			for current.next != nil {
				current = current.next
			}
			current.next = node
		}
		list.size++
	}
}

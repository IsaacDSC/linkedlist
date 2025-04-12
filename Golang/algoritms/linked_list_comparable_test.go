package algoritms

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// go test -v -run Resource ./algoritms
// go test -v -run Optimizations ./algoritms

// Estrutura para armazenar resultados detalhados de benchmark
type ResourceMetrics struct {
	Duration        time.Duration
	AllocBytes      uint64
	TotalAllocBytes uint64
	Mallocs         uint64
	Frees           uint64
	HeapObjects     uint64
	NumGC           uint32
}

// Executa uma função e mede recursos consumidos
func measureResources(operation string, f func()) ResourceMetrics {
	// Força coleta de lixo antes de começar
	runtime.GC()

	// Captura estatísticas iniciais
	var startStats runtime.MemStats
	runtime.ReadMemStats(&startStats)
	startTime := time.Now()

	// Executa a operação
	f()

	// Captura estatísticas finais
	endTime := time.Since(startTime)

	// Força coleta de lixo para medição precisa
	runtime.GC()

	var endStats runtime.MemStats
	runtime.ReadMemStats(&endStats)

	// Calcula diferenças
	metrics := ResourceMetrics{
		Duration:        endTime,
		AllocBytes:      endStats.Alloc - startStats.Alloc,
		TotalAllocBytes: endStats.TotalAlloc - startStats.TotalAlloc,
		Mallocs:         endStats.Mallocs - startStats.Mallocs,
		Frees:           endStats.Frees - startStats.Frees,
		HeapObjects:     endStats.HeapObjects - startStats.HeapObjects,
		NumGC:           endStats.NumGC - startStats.NumGC,
	}

	return metrics
}

// Imprime um relatório de recursos utilizado
func printResourceReport(testName string, linkedList, sliceMetrics ResourceMetrics) {
	fmt.Printf("\n=== Comparação de Recursos: %s ===\n", testName)
	fmt.Printf("%-15s | %-15s | %-15s | %-15s\n", "Estrutura", "Tempo", "Memória Total", "Alocações")
	fmt.Printf("%-15s | %-15s | %-15s | %-15d\n",
		"LinkedList",
		linkedList.Duration,
		formatBytes(linkedList.TotalAllocBytes),
		linkedList.Mallocs)

	fmt.Printf("%-15s | %-15s | %-15s | %-15d\n",
		"Slice",
		sliceMetrics.Duration,
		formatBytes(sliceMetrics.TotalAllocBytes),
		sliceMetrics.Mallocs)

	// Cálculo da diferença
	timeRatio := float64(linkedList.Duration) / float64(sliceMetrics.Duration)
	memoryRatio := float64(linkedList.TotalAllocBytes) / float64(sliceMetrics.TotalAllocBytes)

	fmt.Printf("\nComparação (LinkedList/Slice):\n")
	fmt.Printf("Tempo: %.2fx %s\n",
		timeRatio,
		comparativeText(timeRatio))

	fmt.Printf("Memória: %.2fx %s\n\n",
		memoryRatio,
		comparativeText(memoryRatio))
}

// Retorna texto comparativo
func comparativeText(ratio float64) string {
	if ratio < 0.95 {
		return "mais eficiente"
	} else if ratio > 1.05 {
		return "menos eficiente"
	} else {
		return "equivalente"
	}
}

// Formata bytes para exibição amigável
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// TestResourceInsert compara recursos para operações de inserção
func TestResourceInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando teste de recursos em modo curto")
	}

	const dataSize = 10000

	// Teste 1: Inserir no início
	linkedListMetrics := measureResources("Inserir no início (LinkedList)", func() {
		list := LinkedList[int]{}
		for i := 0; i < dataSize; i++ {
			list.Unshift(i)
		}
	})

	sliceMetrics := measureResources("Inserir no início (Slice)", func() {
		var slice []int
		for i := 0; i < dataSize; i++ {
			slice = append([]int{i}, slice...)
		}
	})

	printResourceReport("Inserção no Início", linkedListMetrics, sliceMetrics)

	// Teste 2: Inserir no fim
	linkedListMetrics = measureResources("Inserir no fim (LinkedList)", func() {
		list := LinkedList[int]{}
		for i := 0; i < dataSize; i++ {
			list.Append(i)
		}
	})

	sliceMetrics = measureResources("Inserir no fim (Slice)", func() {
		var slice []int
		for i := 0; i < dataSize; i++ {
			slice = append(slice, i)
		}
	})

	printResourceReport("Inserção no Final", linkedListMetrics, sliceMetrics)

	// Teste 3: Inserir no meio
	linkedListMetrics = measureResources("Inserir no meio (LinkedList)", func() {
		list := prepareLinkedList(dataSize)
		middle := dataSize / 2
		for i := 0; i < 100; i++ {
			list.AppendOnIndex(i, middle)
		}
	})

	sliceMetrics = measureResources("Inserir no meio (Slice)", func() {
		slice := make([]int, dataSize)
		for i := 0; i < dataSize; i++ {
			slice[i] = i
		}
		middle := len(slice) / 2
		for i := 0; i < 100; i++ {
			slice = append(slice[:middle], append([]int{i}, slice[middle:]...)...)
		}
	})

	printResourceReport("Inserção no Meio", linkedListMetrics, sliceMetrics)
}

// TestResourceSearch compara recursos para operações de busca
func TestResourceSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando teste de recursos em modo curto")
	}

	const dataSize = 10000

	// Teste 1: Busca por valor
	linkedListMetrics := measureResources("Busca por valor (LinkedList)", func() {
		list := prepareLinkedList(dataSize)
		for i := 0; i < 1000; i++ {
			list.Search(dataSize - 10) // Valor perto do final
		}
	})

	sliceMetrics := measureResources("Busca por valor (Slice)", func() {
		slice := make([]int, dataSize)
		for i := 0; i < dataSize; i++ {
			slice[i] = i
		}
		target := dataSize - 10
		for i := 0; i < 1000; i++ {
			for j, v := range slice {
				if v == target {
					_ = j
					break
				}
			}
		}
	})

	printResourceReport("Busca por Valor", linkedListMetrics, sliceMetrics)

	// Teste 2: Busca por índice
	linkedListMetrics = measureResources("Busca por índice (LinkedList)", func() {
		list := prepareLinkedList(dataSize)
		middle := dataSize / 2
		for i := 0; i < 10000; i++ {
			_, _ = list.FindByIndex(middle)
		}
	})

	sliceMetrics = measureResources("Busca por índice (Slice)", func() {
		slice := make([]int, dataSize)
		for i := 0; i < dataSize; i++ {
			slice[i] = i
		}
		middle := dataSize / 2
		for i := 0; i < 10000; i++ {
			_ = slice[middle]
		}
	})

	printResourceReport("Busca por Índice", linkedListMetrics, sliceMetrics)
}

// TestResourceDelete compara recursos para operações de remoção
func TestResourceDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando teste de recursos em modo curto")
	}

	const dataSize = 10000

	// Teste 1: Remover do início
	linkedListMetrics := measureResources("Remover do início (LinkedList)", func() {
		list := prepareLinkedList(dataSize)
		for i := 0; i < 1000; i++ {
			if list.Size() > 0 {
				firstVal, _ := list.FindByIndex(0)
				list.Delete(firstVal)
			}
		}
	})

	sliceMetrics := measureResources("Remover do início (Slice)", func() {
		slice := make([]int, dataSize)
		for i := 0; i < dataSize; i++ {
			slice[i] = i
		}
		for i := 0; i < 1000; i++ {
			if len(slice) > 0 {
				slice = slice[1:]
			}
		}
	})

	printResourceReport("Remoção do Início", linkedListMetrics, sliceMetrics)

	// Teste 2: Remover do meio
	linkedListMetrics = measureResources("Remover do meio (LinkedList)", func() {
		list := prepareLinkedList(dataSize)
		for i := 0; i < 1000; i++ {
			if list.Size() > 0 {
				middle := list.Size() / 2
				middleVal, _ := list.FindByIndex(middle)
				list.Delete(middleVal)
			}
		}
	})

	sliceMetrics = measureResources("Remover do meio (Slice)", func() {
		slice := make([]int, dataSize)
		for i := 0; i < dataSize; i++ {
			slice[i] = i
		}
		for i := 0; i < 1000; i++ {
			if len(slice) > 0 {
				middle := len(slice) / 2
				slice = append(slice[:middle], slice[middle+1:]...)
			}
		}
	})

	printResourceReport("Remoção do Meio", linkedListMetrics, sliceMetrics)
}

// TestOptimizationsImpact mede o impacto das otimizações
func TestOptimizationsImpact(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando teste de otimizações em modo curto")
	}

	const dataSize = 10000

	// Versão otimizada (com tail)
	optimizedMetrics := measureResources("Append Otimizado", func() {
		list := LinkedList[int]{}
		for i := 0; i < dataSize; i++ {
			list.Append(i)
		}
	})

	// Versão não otimizada (sem tail)
	nonOptimizedMetrics := measureResources("Append Não Otimizado", func() {
		list := LinkedList[int]{}
		for i := 0; i < dataSize; i++ {
			// Simular o comportamento sem tail
			node := &Node[int]{i, nil}
			if list.head == nil {
				list.head = node
			} else {
				current := list.head
				for current.next != nil {
					current = current.next
				}
				current.next = node
			}
			list.size++
		}
	})

	fmt.Printf("\n=== Impacto das Otimizações ===\n")
	fmt.Printf("%-20s | %-15s | %-15s | %-15s\n", "Implementação", "Tempo", "Memória Total", "Alocações")
	fmt.Printf("%-20s | %-15s | %-15s | %-15d\n",
		"Com tail",
		optimizedMetrics.Duration,
		formatBytes(optimizedMetrics.TotalAllocBytes),
		optimizedMetrics.Mallocs)

	fmt.Printf("%-20s | %-15s | %-15s | %-15d\n",
		"Sem tail",
		nonOptimizedMetrics.Duration,
		formatBytes(nonOptimizedMetrics.TotalAllocBytes),
		nonOptimizedMetrics.Mallocs)

	// Cálculo da melhoria
	timeImprovement := float64(nonOptimizedMetrics.Duration) / float64(optimizedMetrics.Duration)

	fmt.Printf("\nMelhoria de performance: %.2fx mais rápido com otimizações\n", timeImprovement)
}

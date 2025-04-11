package main

import (
	"linkedlist/Golang/algoritms"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

// Variável global para controlar se os logs de recursos serão exibidos
var shouldReportResources = false

func reportResourceUsage(b *testing.B) {
	if !shouldReportResources {
		return
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	b.Logf("Memória Alocada: %d bytes", memStats.Alloc)
	b.Logf("Número de Goroutines: %d", runtime.NumGoroutine())
}

// TestMain será executado antes dos benchmarks para configurar o profiling
func TestMain(m *testing.M) {
	// Criar um único arquivo de perfil de CPU para todos os benchmarks
	f, err := os.Create("cpu.prof")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		os.Exit(1)
	}
	defer pprof.StopCPUProfile()

	// Executar todos os testes e benchmarks
	code := m.Run()
	os.Exit(code)
}

func BenchmarkLinkedListWrite(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceWrite(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListRead(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Search(i % 1000)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceRead(b *testing.B) {
	var slice []int
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slice[i%1000]
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListUnshift(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Unshift(i)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceUnshift(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append([]int{i}, slice...)
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListAppend(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceAppend(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListAppendOnIndex(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.AppendOnIndex(i, i%1000)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceAppendOnIndex(b *testing.B) {
	var slice []int
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := i % (len(slice) + 1)
		slice = append(slice[:index], append([]int{i}, slice[index:]...)...)
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListSearch(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Search(i % 1000)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceSearch(b *testing.B) {
	var slice []int
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			if v == i%1000 {
				break
			}
		}
	}
	reportResourceUsage(b)
}

func BenchmarkLinkedListDelete(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Delete(i % 1000)
	}
	reportResourceUsage(b)
}

func BenchmarkSliceDelete(b *testing.B) {
	var slice []int
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if len(slice) == 0 {
			continue
		}
		index := i % len(slice)
		slice = append(slice[:index], slice[index+1:]...)
	}
	reportResourceUsage(b)
}

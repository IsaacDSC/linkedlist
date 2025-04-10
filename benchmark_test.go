package main

import (
	"linkedlist/algoritms"
	"runtime"
	"testing"
)

func reportMemoryUsage(b *testing.B) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	b.Logf("Memória Alocada: %d bytes", memStats.Alloc)
}

func BenchmarkLinkedListWrite(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
	reportMemoryUsage(b)
}

func BenchmarkSliceWrite(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
}

func BenchmarkLinkedListUnshift(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Unshift(i)
	}
	reportMemoryUsage(b)
}

func BenchmarkSliceUnshift(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append([]int{i}, slice...)
	}
	reportMemoryUsage(b)
}

func BenchmarkLinkedListAppend(b *testing.B) {
	list := algoritms.LinkedList[int]{}
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
	reportMemoryUsage(b)
}

func BenchmarkSliceAppend(b *testing.B) {
	var slice []int
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
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
	reportMemoryUsage(b)
}

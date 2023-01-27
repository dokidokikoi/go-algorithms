package sorting

import (
	"go-algorithms/tools"
	"sort"
	"testing"
)

var funcs = []struct {
	name string
	f    Sort
}{
	{"bubble", BubbleSort},
	{"Insert", InsertionSort},
	{"Select", SelectionSort},
	{"Heap", HeapSort},
	{"Shell", ShellSort},
	{"Count", CountSort},
	{"Bucket", BucketSort},
}

func TestSort(t *testing.T) {
	for _, tt := range funcs {
		t.Run(tt.name, func(t *testing.T) {
			arr := tools.RandArray(10)

			tt.f(arr)

			if !sort.IntsAreSorted(arr) {
				t.Errorf("%v is not sorted", arr)
			}
			t.Logf("\narr: %v", arr)
		})
	}
}

func BenchmarkSort(b *testing.B) {
	for _, tt := range funcs {
		arrs := make([][]int, 0)

		for n := 0; n < b.N; n++ {
			arrs = append(arrs, tools.RandArray(10))
		}

		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < len(arrs); n++ {
				tt.f(arrs[n])
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	arr := tools.RandArray(10)
	t.Logf("\nInitial array: %v", arr)
	t.Logf("\nSorted array: %v", QuickSort(arr))
}

func TestMergeSort(t *testing.T) {
	arr := tools.RandArray(10)
	t.Logf("\nInitial array: %v", arr)
	t.Logf("\nSorted array: %v", MergeSort(arr))
}

func TestRadixSort(t *testing.T) {
	var arr = []int32{421, 15, -175, 90, -2, 214, -52, -166}
	t.Logf("\nInitial array: %v", arr)
	RadixSort(arr)
	t.Logf("\nSorted array: %v", arr)
}

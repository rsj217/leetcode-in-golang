package sort

import "testing"

func TestAlgoRandomPerformance(t *testing.T) {
	algos := []Algo{
		{"SelectSort", SelectSort},
		{"InsertSort", InsertSort},
		{"QuickSort", QuickSort},
		{"MergeSort", MergeSort},
		{"HeapSort", HeapSort},
	}
	AlgoRandomPerformance(algos, 100000, 0, 10000)
}

func TestAlgoNearlyOrderPerformance(t *testing.T) {
	algos := []Algo{
		{"InsertSort", InsertSort},
		{"QuickSort", QuickSort},
		{"MergeSort", MergeSort},
		{"HeapSort", HeapSort},
	}
	AlgoNearlyOrderPerformance(algos, 100000, 10)
}

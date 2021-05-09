package datastruct

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range nums {
		h.Push(v)
	}
	fmt.Println(h)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	h.Heapify(nums)

	fmt.Println(h)

	heapPrint(nums)
}

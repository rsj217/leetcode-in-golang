package sort

func heapSort(nums []int, lo, hi int) {
	heapify(nums, lo, hi)
	for i := hi - 1; lo < i; i-- {
		nums[lo], nums[i] = nums[i], nums[lo]
		shiftDown(nums, lo, i, lo)
	}
}

func heapify(nums []int, lo, hi int) {
	for i := (hi - 1) / 2; lo <= i; i-- {
		shiftDown(nums, lo, hi, i)
	}
}

func shiftDown(nums []int, lo, hi, k int) {
	for 2*k+1 < hi {
		j := 2*k + 1
		if j+1 < hi && nums[j] < nums[j+1] {
			j++
		}
		if nums[j] <= nums[k] {
			break
		}
		nums[j], nums[k] = nums[k], nums[j]
	}
}

func HeapSort(nums []int) {
	heapSort(nums, 0, len(nums))
}

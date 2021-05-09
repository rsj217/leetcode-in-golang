package datastruct

import (
	"fmt"
	"math"
	"strconv"
)

type MaxHeap []int

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) GetSize() int {
	return len(*h)
}

func (h *MaxHeap) Push(item int) {
	*h = append(*h, item)
	shiftUp(*h, h.GetSize()-1)
}

func (h *MaxHeap) Pop() int {
	ans := 0
	size := h.GetSize()
	if 0 < size {
		(*h)[0], (*h)[size-1] = (*h)[size-1], (*h)[0]
		ans = (*h)[size-1]
		*h = (*h)[:size-1]
		shiftDown(*h, 0)
	}
	return ans
}

func (h *MaxHeap) Heapify(nums []int) {
	heapify(nums)
	*h = nums
}

func heapify(nums []int) {
	for k := len(nums) / 2; 0 <= k; k-- {
		shiftDown(nums, k)
	}
}

func shiftUp(nums []int, k int) {
	for 0 < k && nums[(k-1)/2] < nums[k] {
		nums[(k-1)/2], nums[k] = nums[k], nums[(k-1)/2]
		k = (k - 1) / 2
	}
}

func shiftDown(nums []int, k int) {

	for 2*k+1 < len(nums) {
		j := 2*k + 1
		if j+1 < len(nums) && nums[j] < nums[j+1] {
			j = j + 1
		}
		if nums[j] <= nums[k] {
			break
		}
		nums[k], nums[j] = nums[j], nums[k]
		k = j
	}
}

// 以树状打印整个堆结构
func heapPrint(nums []int) {

	n := len(nums)
	size := len(nums)

	maxLevel := 0
	numberPerLevel := 1
	for n > 0 {
		maxLevel += 1
		n -= numberPerLevel
		numberPerLevel *= 2
	}

	maxLevelNumber := int64(math.Pow(2, float64(maxLevel-1)))
	curTreeMaxLevelNumber := maxLevelNumber
	index := 1
	for level := 0; level < maxLevel; level++ {

		line1 := fmt.Sprintf("%*s\n", maxLevelNumber*3-1, "")

		curLevelNumber := int64(math.Min(float64(size)-math.Pow(2, float64(level))+1, math.Pow(2, float64(level))))
		isLeft := true
		indexCurLevel := int64(0)
		for indexCurLevel < curLevelNumber {
			line1 = putNumberInLine(nums[index-1], line1, int(indexCurLevel), int(curTreeMaxLevelNumber*3-1), isLeft)
			isLeft = !isLeft
			index++
			indexCurLevel++
		}
		fmt.Printf(line1)

		if level == maxLevel-1 {
			break
		}

		line2 := fmt.Sprintf("%*s\n", maxLevelNumber*3-1, "")
		for indexCurLevel = 0; indexCurLevel < curLevelNumber; indexCurLevel++ {
			line2 = putBranchInLine(line2, int(indexCurLevel), int(curTreeMaxLevelNumber*3-1))
		}
		fmt.Printf(line2)
		curTreeMaxLevelNumber /= 2
	}
}

func putNumberInLine(num int, line string, indexCurLevel int, curTreeWidth int, isLeft bool) string {

	subTreeWidth := (curTreeWidth - 1) / 2
	offset := indexCurLevel*(curTreeWidth+1) + subTreeWidth
	if num >= 10 {

		line = line[0:offset+0] + strconv.Itoa(num) + line[offset+2:]
	} else {
		if isLeft {
			line = line[0:offset+0] + strconv.Itoa(num) + line[offset+1:]
		} else {
			line = line[0:offset+1] + strconv.Itoa(num) + line[offset+2:]
		}
	}
	return line
}

func putBranchInLine(line string, indexCurLevel, curTreeWidth int) string {

	subTreeWidth := (curTreeWidth - 1) / 2
	subSubTreeWidth := (subTreeWidth - 1) / 2
	offsetLeft := indexCurLevel*(curTreeWidth+1) + subSubTreeWidth
	offsetRight := indexCurLevel*(curTreeWidth+1) + subTreeWidth + 1 + subSubTreeWidth

	line = line[0:offsetLeft+1] + "/" + line[offsetLeft+2:]
	line = line[0:offsetRight] + "\\" + line[offsetRight+1:]
	return line
}

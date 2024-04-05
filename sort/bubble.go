package sort

func BubbleSort(nums []int) []int {
	copiedNums := CopySlice(nums)

	lenNums := len(copiedNums)

	for i := 0; i < lenNums; i++ {
		for j := 0; j < lenNums - 1 - i; j++ {
			if copiedNums[j] > copiedNums[j+1] {
				copiedNums[j], copiedNums[j+1] = copiedNums[j+1], copiedNums[j]
			}
		}
	}

	return copiedNums
}

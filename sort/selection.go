package sort

func SelectionSort(nums []int) []int {
	copiedNums := CopySlice(nums)

	lenNums := len(copiedNums)

	for i := 0; i < lenNums; i++ {
		index := i

		for j := i + 1; j < lenNums; j++ {
			if copiedNums[index] > copiedNums[j] {
				index = j
			}
		}
		copiedNums[index], copiedNums[i] = copiedNums[i], copiedNums[index]
	}

	return copiedNums
}

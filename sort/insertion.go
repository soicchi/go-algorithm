package sort

func InsertionSort(nums []int) []int {
	copiedNums := CopySlice(nums)

	for i := 1; i < len(copiedNums); i++ {
		tmp := copiedNums[i]
		j := i - 1
		for j >= 0 && copiedNums[j] > tmp {
			copiedNums[j+1] = copiedNums[j]
			j--
		}
		copiedNums[j+1] = tmp 
	}

	return copiedNums
}

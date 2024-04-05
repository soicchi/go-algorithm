package sort

func BubbleSort(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		for j := 0; j < lenNums - 1 - i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

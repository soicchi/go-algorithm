package sort

func CopySlice(base []int) []int {
	copied := make([]int, len(base))
	copy(copied, base)

	return copied
}

package main

import (
	"fmt"

	"algorithm/sort"
)

var baseNums = []int{2, 5, 1, 8, 7, 3}

func main() {
	// bubble sort
	numsForBubble := make([]int, len(baseNums))
	copy(numsForBubble, baseNums)

	sort.BubbleSort(numsForBubble)
	fmt.Println("Bubble sort:", numsForBubble)
}

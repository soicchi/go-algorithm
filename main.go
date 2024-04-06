package main

import (
	"fmt"

	"algorithm/sort"
)

var baseNums = []int{2, 5, 1, 8, 7, 3}

func main() {
	// bubble sort
	bubbleSlice := sort.BubbleSort(baseNums)
	fmt.Println("Bubble sort:", bubbleSlice)

	// selection sort
	selectionSlice := sort.SelectionSort(baseNums)
	fmt.Println("Selection sort:", selectionSlice)

	// insertion sort
	insertionSlice := sort.InsertionSort(baseNums)
	fmt.Println("Insertion sort:", insertionSlice)
}

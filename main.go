package main

import (
	"sort"

	. "./comm"
	"./msort"
)

func main() {
	sortDemo()
}

func sortDemo() {
	msort.Version()
	nums := []int{4, 6, 3, 1, 2, 5}
	PrintIntSlice("before sort", nums)
	//msort.InsertSort(nums)
	//msort.BubbleSort(nums)
	//msort.SelectionSort(nums)
	//msort.ShellSort(nums)
	//msort.QuickSort(nums)
	//msort.QuickSort2(nums)
	//msort.HeapSort(nums)
	msort.MergeSort(nums)
	PrintIntSlice("after sort", nums)

	nums = []int{4, 6, 3, 1, 2}
	sort.Ints(nums)
	PrintIntSlice("std sort", nums)
}

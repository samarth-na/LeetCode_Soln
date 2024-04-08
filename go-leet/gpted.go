package main

import (
	"fmt"
	"sort"
)

func SplitArray(nums []int) bool {
	sort.Ints(nums)
	fmt.Println(nums)

	var ar1, ar2 []int
	for i := 0; i < len(nums); i += 2 {
		ar1 = append(ar1, nums[i])
		if i+1 < len(nums) { // Check if i+1 is within bounds
			ar2 = append(ar2, nums[i+1])
		}
	}
	fmt.Println(ar1, ar2)

	for i := range ar1 {
		if ar1[i] != ar2[i] {
			return true
		}
	}

	return false
}

func GPTed() {
	ints := []int{6, 1, 3, 1, 1, 8, 9, 2}
	fmt.Println(SplitArray(ints))
}

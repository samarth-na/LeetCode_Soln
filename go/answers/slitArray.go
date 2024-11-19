package answers

import (
	"fmt"
	"sort"
)

// [8,1,2,2,3,4,9,2]
func SplitArray(nums []int) bool {
	var ar1, ar2 []int
	sort.Ints(nums)
	fmt.Print(nums, "\n")

	var i int
	for i < len(nums) {
		ar1 = append(ar1, nums[i])
		ar2 = append(ar2, nums[i+1])
		i += 2
	}
	fmt.Print(ar1, ar2)

	var istrue bool
	var truelist []bool
	for i := range ar1 {
		istrue = ar1[i] != ar2[i]
		truelist = append(truelist, istrue)
	}
	fmt.Print(truelist)
	for _, b := range truelist {
		if b {
			return true
		}
	}

	return false
}

// problem 268
package answers

import (
	"fmt"
	"sort"
)

// {3,0,1} =1
func MissingNo(nums []int) int {
	sort.Ints(nums)
	assenArray := nums

	fmt.Print(assenArray, "\n")

	var missing int
	for i, v := range assenArray {
		if i+1 < len(assenArray) {
			if assenArray[i+1]-v != 1 {
				missing = assenArray[i+1] - 1
			}
		}
	}

	return missing
}

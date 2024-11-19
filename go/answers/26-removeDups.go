package problems

// 1, 1, 2, 3, 4, 5, 6, 6

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIndex := 1
	prevValue := nums[0]
	for v := range nums[1:] {
		if v != prevValue {
			nums[uniqueIndex] = v
			uniqueIndex++
			prevValue = v
		}
	}
	return uniqueIndex
}

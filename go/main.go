package main

import (
	"fmt"
	"go-leet/answers"
)

func main() {
	ints := []int{6, 1, 3, 1, 1, 8, 9, 2}
	// fmt.Print(ans)
	ans := answers.SplitArray(ints)
	fmt.Print("\n \n ", ans)
}

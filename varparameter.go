package main

import (
	"fmt"
)

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	x := min(1, 3, 2, 4)
	fmt.Println("min1:", x)

	slice := []int{7, 3, 4, 9}
	x = min(slice...)
	fmt.Println(x)
}

package main

import (
	"fmt"
)

/* Time Complexity: O(N)
Space Complexity: O(N) */
func rotate(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	R := len(A) - K%len(A)
	return append(A[R:], A[:R]...)
}

func main() {
	rotatedArr := rotate([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(rotatedArr)
}

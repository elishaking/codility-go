package main

import (
	"fmt"
)

func rotate(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	R := len(A) - K%len(A)
	A = append(A[R:], A[:R]...)

	return A
}

func main() {
	rotatedArr := rotate([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(rotatedArr)
}

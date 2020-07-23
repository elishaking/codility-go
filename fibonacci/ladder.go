package main

import (
	"fmt"
	"math"
)

/* Time Complexity: O(N)
Space Complexity: O(N) */
func ladder(A []int, B []int) []int {
	ways := []int{1, 1}
	maxA := 0
	maxB := 0

	for i := 0; i < len(A); i++ {
		if A[i] > maxA {
			maxA = A[i]
		}
		if B[i] > maxB {
			maxB = B[i]
		}
	}

	for i := 1; i <= maxA; {
		i++
		ways = append(ways, (ways[i-1]+ways[i-2])%int(math.Pow(2, float64(maxB))))
	}

	var R []int
	for i := 0; i < len(A); i++ {
		R = append(R, ways[A[i]]&int(math.Pow(2, float64(B[i]))-1))
	}

	return R
}

func main() {
	R := ladder([]int{4, 4, 5, 5, 1}, []int{3, 2, 4, 3, 1})
	fmt.Println(R)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	// get m and n, sanitize it for errors
	if len(os.Args) != 3 {
		fmt.Printf("There should be 2 arguments")
		return
	}

	mStr, nStr := os.Args[1], os.Args[2]
	m, err := strconv.Atoi(mStr)
	if err != nil || m < 1 {
		fmt.Printf("1st argument should be a positive integer")
		return
	}
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		fmt.Printf("2nd argument should be a positive integer")
		return
	}

	// build dynamic table (+1 to consider empty string)
	A := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		A[i] = make([]int, n+1)
		A[i][0] = 1 // base case
	}
	for j := 0; j < n+1; j++ {
		A[0][j] = 1 // base case
	}

	// build rest of dynamic table
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			A[i][j] = A[i-1][j] + A[i][j-1] + A[i-1][j-1] // core recurrence equation
		}
	}

	// print solution (which is at bottom-right corner of the dynamic table)
	fmt.Printf("The toal number of possible alignments from strings of length %d and %d is .. %d (took %s)", m, n, A[m][n], time.Since(start))
}

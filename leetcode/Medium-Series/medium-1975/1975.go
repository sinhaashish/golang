// You are given an n x n integer matrix. You can do the following operation any number of times:

// Choose any two adjacent elements of matrix and multiply each of them by -1.
// Two elements are considered adjacent if and only if they share a border.

// Your goal is to maximize the summation of the matrix's elements. Return the maximum sum of the matrix's elements using the operation mentioned above.
package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	count := 0
	var sum int64
	min := math.MaxInt
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] < 0 {
				count++
			}
			min = int(math.Min(float64(min), float64(math.Abs(float64(matrix[i][j])))))
			sum += int64(math.Abs(float64(matrix[i][j])))
		}
	}
	if count%2 == 0 {
		return sum
	}
	return sum - int64(2*min)

}

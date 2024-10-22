package main

import "fmt"

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	columns := len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	maxSide := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1 // 第一行和第一列
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	fmt.Print(dp)
	return maxSide * maxSide // 返回最大正方形的面积
}
func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	result := maximalSquare(matrix)
	fmt.Print(result)
}

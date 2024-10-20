package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j] //核心
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	// 查找最后一行的最小值
	minPathSum := dp[rows-1][0]
	for j := 1; j < rows; j++ {
		minPathSum = min(minPathSum, dp[rows-1][j])
	}
	return minPathSum
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	result := minimumTotal(triangle)
	fmt.Print(result)
}

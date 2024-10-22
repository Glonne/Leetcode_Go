package main

import "fmt"

func minFallingPathSum(matrix [][]int) int {
	rows := len(matrix)
	column := len(matrix)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, column)
	}
	// 边缘第一行初始化
	for j := 0; j < column; j++ {
		dp[0][j] = matrix[0][j]
	}
	// 第二行开始计算dp
	for i := 1; i < rows; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + matrix[i][0] // dp的第一列靠 dp[i-1][0] 与 dp [i-1][1]
		for j := 1; j < column-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
		}
		dp[i][column-1] = min(dp[i-1][column-1], dp[i-1][column-2]) + matrix[i][column-1] // dp的第一列靠 dp[i-1][column-1] 与 dp[i-1][column-2]
	}
	minnum := dp[rows-1][0]
	for j := 1; j < column; j++ {
		if dp[rows-1][j] < minnum {
			minnum = dp[rows-1][j]
		} else {
			continue
		}
	}
	return minnum
}
func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	result := minFallingPathSum(matrix)
	fmt.Print(result)
}

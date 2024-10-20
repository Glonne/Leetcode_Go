package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			break
		}
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}
	fmt.Print(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Print(dp)
	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	result := uniquePathsWithObstacles(grid)
	fmt.Println(result)
}

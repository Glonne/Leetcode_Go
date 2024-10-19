package main

import "fmt"

func minPathSum(grid [][]int) int {
	rows := len(grid)
	colunm := len(grid[0])
	fmt.Printf("row = %d,column = %d\n", rows, colunm)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, colunm)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < colunm; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < colunm; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	//fmt.Printf("dp[0][0] = %d \n", dp[0][0])
	//fmt.Printf("dp[1][1] = %d \n", dp[1][1])
	//fmt.Printf("dp[2][2] = %d \n", dp[2][2])
	return dp[rows-1][colunm-1]
}
func main() {
	grid := make([][]int, 0)
	grid = append(grid, []int{1, 3, 1})
	grid = append(grid, []int{1, 5, 1})
	grid = append(grid, []int{4, 2, 1})
	fmt.Print(minPathSum(grid))
	//fmt.Println(grid)
}

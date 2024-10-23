// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
package main

import "fmt"

func main() {
	s1 := "sea"
	s2 := "eat"
	result := minimumDeleteSum(s1, s2)
	fmt.Print(result)
}
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := range dp[0] {
		if j > 0 {
			dp[0][j] = dp[0][j-1] + int(s2[j-1])
		}
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(c1), dp[i+1][j]+int(c2))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

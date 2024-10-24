package main

import "fmt"

func main() {
	num := minDistance("waw", "wowd")
	fmt.Print(num)
}
func minDistance(word1 string, word2 string) int {
	// 增加a
	// 增加b
	// 修改a
	// dp[i][j] 代表 word1 的前i位置，和word2的前j位置更改步数
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

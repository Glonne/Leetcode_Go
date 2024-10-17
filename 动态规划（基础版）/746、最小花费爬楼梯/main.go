// 0 ~ n-1台阶
// const【】数组存放花费
// 第n个台阶为顶层 假如为dp[],dp[i]代表登顶i最少消费
// dp[0] = 0 dp [1]= 1  dp[2] = min {dp[1],dp[0]}  dp[3]=min{dp[1],dp[2]}
package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 15, 20}
	fmt.Printf("The cost is :%d \n", minCostClimbingStairs(nums))
}

// 类似组合数学理论
// n个房屋之间的最高金额组合 f(n) = max(f(n-1),f(n-2)+cost[n])
package main

import (
	"fmt"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 15, 20}
	fmt.Printf("The cost is :%d \n", rob(nums))
}

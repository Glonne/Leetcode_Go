// 给你一个整数数组 nums ，你可以对它进行一些操作。

// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
package main

import "fmt"

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	score := make([]int, maxVal+1)
	dp := make([]int, maxVal+1)
	for i := 0; i < len(nums); i++ {
		score[nums[i]] += nums[i] //获得相同元素得分
	}

	dp[0] = score[0]
	dp[1] = max(score[0], score[1])
	for i := 2; i < len(score); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+score[i])
	}
	return dp[len(score)-1]

}

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	fmt.Printf("The cost is :%d \n", deleteAndEarn(nums))
}

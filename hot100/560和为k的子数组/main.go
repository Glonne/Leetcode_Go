package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0
	// 1 2 3
	// preSum 1  3  6
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}

func main() {
	// 示例测试用例
	nums := []int{1, 2, 3}
	k := 3

	result := subarraySum(nums, k)
	fmt.Printf("Number of subarrays with sum %d: %d\n", k, result)
}

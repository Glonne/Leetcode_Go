package main

import "fmt"

func main() {
	nums1 := []int{1, 0, 0, 0, 4, 5, 6}
	moveZeroes(nums1)
	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1[i])
	}
}
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 双指针不一定是一个往左一个往右
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

package main

import "fmt"

// 你需要自己实现这个函数
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0) //存放result结果。
	queue := make([]int, 0)  //存放最大值的队列的索引

	for i, v := range nums {
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] < v {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}

// 打印切片
func printSlice(slice []int) {
	fmt.Print("[")
	for i, v := range slice {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

func main() {
	// 测试用例 1
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Println("输入：", nums1, "窗口大小：", k1)
	fmt.Print("输出：")
	printSlice(maxSlidingWindow(nums1, k1)) // 期望输出 [3,3,5,5,6,7]

	// 测试用例 2
	nums2 := []int{1}
	k2 := 1
	fmt.Println("输入：", nums2, "窗口大小：", k2)
	fmt.Print("输出：")
	printSlice(maxSlidingWindow(nums2, k2)) // 期望输出 [1]

	// 测试用例 3
	nums3 := []int{9, 10, 9, -7, -4, -8, 2, -6}
	k3 := 5
	fmt.Println("输入：", nums3, "窗口大小：", k3)
	fmt.Print("输出：")
	printSlice(maxSlidingWindow(nums3, k3)) // 期望输出 [10,10,9,2]
}

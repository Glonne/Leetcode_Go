// 两数之和
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	nums1 := twoSum1(nums, target)
	for i := 0; i < 2; i++ {
		fmt.Printf("nums= %d \n", nums1[i])
	}

}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。
func twoSum1(nums []int, target int) []int {
	HashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := HashTable[target-x]; ok {
			return []int{p, i}
		}
		HashTable[x] = i
	}
	return nil
}

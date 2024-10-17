// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 类似加法定理
package main

import (
	"fmt"
)

// ==================================超时==================================
func climbStairs(n int) int {
	if n < 1 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var num int = climbStairs(n-1) + climbStairs(n-2)
	return num
	//return climbStairs(n-1) + climbStairs(n-2)

}

// ==============================直接利用for循环===========================
func climbStairs1(n int) int {
	p, q := 1, 2
	if n < 1 {
		return 1
	}
	if n == 1 {
		return p
	} else if n == 2 {
		return q
	} else {
		var r int
		for i := 3; i <= n; i++ {
			r = q + p
			p = q
			q = r
		}
		return r
	}
}
func main() {

	fmt.Printf("The number is %d ", climbStairs1(5))

}

// 泰波那契序列 Tn 定义如下：

// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

// 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

package main

import "fmt"

func tribonacci(n int) int {
	var num int
	p, q, r := 0, 1, 1
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		for i := 3; i <= n; i++ {
			num = p + q + r
			p = q
			q = r
			r = num
		}
		return num
	}

}
func main() {
	fmt.Printf("The number is %d\n", tribonacci(4))
	fmt.Printf("The number is %d\n", tribonacci(25))
}

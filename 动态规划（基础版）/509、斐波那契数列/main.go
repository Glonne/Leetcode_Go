package main

import "fmt"

func main() {
	fmt.Printf("The number is :%d", fib(3))
}
func fib(n int) int {
	p, q := 0, 1
	var r int
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}
	return r

}

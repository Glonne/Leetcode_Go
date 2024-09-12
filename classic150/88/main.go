// 合并两个数组
package main
import "fmt"
func main(){
	map1 := map[int]string{
		1:"abc",
	}
	if _,exist:=map1[1];exist{
		fmt.Println("有")
		fmt.Println(exist)
	}
}

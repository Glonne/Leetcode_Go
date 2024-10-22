package main

import (
	"fmt"
)

func main() {
	s := "apple"
	worddict := make([]string, 0)
	worddict = append(worddict, "apple", "pen")
	fmt.Println(wordBreak(s, worddict))
	fmt.Print(s)
	fmt.Print(worddict)
}

// dp[i] 代表 string的前i个是否可以被检索
// dp[i] = dp[j] && check (j,i-1) j是0-i-1
func check(s string, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return true
		}
	}
	return false
}
func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	for i := 1; i < len(s)+1; i++ {
		dp[i] = false
	}
	dp[0] = true
	if len(s) <= 1 {
		return check(s, wordDict)
	} else {
		for i := 1; i <= len(s); i++ {
			for j := 0; j < i; j++ {
				if dp[j] && check(s[j:i], wordDict) {
					dp[i] = true
					fmt.Print(dp[i])
				}
			}
		}
	}
	return dp[len(s)]
}

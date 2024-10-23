package main

import "fmt"

// 给你一个字符串 s，找到 s 中最长的回文子串
// s = "babad" bab aba
func main() {
	fmt.Println(longestPalindrome("abac"))
}

// dp[i][j]判断是否为s[i:j]为回文字符串
// dp[i][j] 如果dp[i][j] 是，而且dp[i+1][j-1] 是，且string[i] = string[j]
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	// 创建一个二维数组,i = j 是回文字符串，下面只需要填充i<j的部分
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}

	maxlength := 1
	start := 0

	for j := 1; j < length; j++ { // 右边界 j
		for i := 0; i < j; i++ { // 左边界 i
			if s[i] == s[j] { // 如果字符相等
				if j-i == 1 || dp[i+1][j-1] == 1 { // 子串长度为 2 或 内部是回文
					dp[i][j] = 1
					if j-i+1 > maxlength { // 更新最大长度
						maxlength = j - i + 1
						start = i
					}
				}
			}
		}
	}
	//fmt.Print(dp)
	return s[start : start+maxlength]
}

package main

import "fmt"

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//s 由英文字母、数字、符号和空格组成

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var windows [256]int

	l, r := 0, 0
	var maxNum int
	//循环停止条件,应该是小于，如果等于，就会越界
	for r < len(s) {
		//s[r]不存在ss中
		windows[s[r]]++
		for windows[s[r]] > 1 {
			windows[s[l]]--
			l++
		}
		maxNum = max(maxNum, r-l+1)
		r++
	}
	return maxNum
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "abcabcbb"
	v := lengthOfLongestSubstring(s)
	fmt.Println(v)
}

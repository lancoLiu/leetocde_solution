package main

import "strings"

//给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
func longestSubstring(s string, k int) (ans int) {
	if len(s) == 0 {
		return
	}
	var cnt [26]int

	for _, v := range s {
		cnt[v-'a']++
	}
	var split byte
	for i, c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//输入：s = "ababbc", k = 2
//输出：5
//解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
func main() {

}

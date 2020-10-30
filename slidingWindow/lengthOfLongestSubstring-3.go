package slidingWindow

/**
定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res, left, right := 0, 0, -1
	var windows [256]int
	for left < len(s) {
		if right+1 < len(s) && windows[s[right+1]-'a'] == 0 {
			windows[s[right+1]-'a']++
			right++
		} else {

			windows[s[left]-'a']--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

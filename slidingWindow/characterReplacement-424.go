package slidingWindow

/**
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意:
字符串长度 和 k 不会超过 104。

示例 1:

输入:
s = "ABAB", k = 2

输出:
4

解释:
用两个'A'替换为两个'B',反之亦然。

*/
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	maxRes := 0
	l := 0
	r := 0
	windows := make(map[uint8]int)
	for ; r < len(s); r++ {
		windows[s[r]-'A']++
		maxRes = max(maxRes, windows[s[r]-'A'])
		if r-l+1 > maxRes+k {
			windows[s[l]-'A']--
			l++
		}
	}
	return r - l
}

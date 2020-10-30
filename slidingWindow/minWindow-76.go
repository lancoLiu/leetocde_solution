package slidingWindow

import "math"

/**
给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。
示例：

输入：S = "ADOBECODEBANC", T = "ABC"
输出："BANC"

*/
func minWindow(s string, t string) string {
	if len(s) == 0 {
		return ""
	}
	//初始化两个map
	windows, need := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	start, lens := 0, math.MaxInt32
	//左闭右开区间
	//valid表示满足need条件的字符个数
	left, right, valid := 0, 0, 0
	//寻找区间[left,right)
	for right < len(s) {
		//往右边移动的字符串
		tempRightStr := s[right]
		right++
		//往窗口加入字符，需要的操作：
		//判断是否是need需要的字符信息
		if _, ok := need[tempRightStr]; ok {
			windows[tempRightStr]++
			if windows[tempRightStr] == need[tempRightStr] {
				valid++
			}
		}

		//什么时候应该停止right移动，暂停扩大，开始移动left缩小窗口
		for valid == len(need) {
			//这一步的操作主要是更新最小覆盖子串
			if right-left < lens {
				start = left
				lens = right - left
			}

			//左移窗口
			tempLeftStr := s[left]
			left++
			//如果在need中，更新数据：winodows窗口数据和valid是否需要相减
			if _, ok := need[tempLeftStr]; ok {
				windows[tempLeftStr]--
				if windows[tempLeftStr] < need[tempLeftStr] {
					valid--
				}
			}
		}
	}
	if lens == math.MaxInt32 {
		return ""
	}
	return s[start : start+lens]
}

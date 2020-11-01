package suanfa

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
*/
func CanPermutePalindrome(s string) bool {

	//taccat tacocat
	//题目可以转为判断是否最多一个字符串是单独的，其他都是成双成对

	var r uint64
	for _, v := range s {
		r ^= 1 << int(v)
	}

	return r&(r-1) == 0
	//
}

func canPermutePalindrome(s string) bool {
	strmap := make(map[rune]int, len(s))
	for _, v := range s {
		strmap[v]++
		if strmap[v]%2 == 0 {
			delete(strmap, v)
		}
	}
	return len(strmap) <= 1
}

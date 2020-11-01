package suanfa

/*
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false


*/

//位运算

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	bits := 0
	for i := 0; i < len(s1); i++ {
		index1 := 1 << (s1[i] - 'a')
		index2 := 1 << (s2[i] - 'a')
		bits += index1
		bits -= index2
	}
	return bits == 0
}

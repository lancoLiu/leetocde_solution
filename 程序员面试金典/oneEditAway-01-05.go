package suanfa

import "math"

/**
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。



示例 1:

输入:
first = "pale"
second = "ple"
输出: True

*/
func oneEditAway(first string, second string) bool {
	//不用考虑
	if first == second {
		return true
	}
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	i, j, k := 0, len(first)-1, len(second)-1
	for i <= j && i <= k && first[i] == second[i] {
		i++
	}
	for j >= 0 && k >= 0 && first[j] == second[k] {
		j--
		k--
	}
	return j-i < 1 && k-i < 1
}

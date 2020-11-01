package suanfa

/**
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true

*/

func isUnique(astr string) bool {
	if len(astr) == 0 {
		return true
	}
	var maps = make(map[int32]bool)
	for _, v := range astr {
		if _, ok := maps[v]; ok {
			return false
		}
		maps[v] = true
	}
	return true

}

//不使用额外结构，位运算，如果所有的字符串都不一样

func isUnique2(astr string) bool {
	bits := 0
	for _, num := range astr {
		if bits|(1<<(num-'a')) != 0 {
			return false
		} else {
			bits = bits | (1 << (num - 'a'))
		}
	}
	return true

}

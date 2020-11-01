package suanfa

import "strconv"

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。

*/

//根据实例是说最近重复的字符串用数字表示

func CompressString(S string) string {
	if len(S) == 0 {
		return S
	}

	//保存最近出现的字符串
	least := string(S[0])
	//保存最近字符串出现的次数
	sum := 0
	res := ""
	for i := 0; i < len(S); i++ {

		if string(S[i]) == least {
			sum++
		} else {
			res += least
			res += strconv.Itoa(sum)
			least = string(S[i])
			sum = 1
		}
	}
	res += least
	res += strconv.Itoa(sum)
	if len(S) <= len(res) {
		return S
	} else {
		return res
	}

}

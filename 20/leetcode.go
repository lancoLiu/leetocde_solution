package main

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

//有效字符串需满足：

//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//test case
//1、}
//2、()
//3、([])

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	//遍历
	for i := 0; i < n; i++ {
		//
		if pairs[s[i]] > 0 {
			//stack里面没有元素，说明都没有左括号了，但是现在要进来一个右括号
			//如果stack有长度，因为stack的进入顺序，所以stack的最后一个元素必须和当前右括号匹配，否则false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			//元素添加进stack中
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

package Backtracking

//https://leetcode-cn.com/problems/generate-parentheses/
/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/
//采用递归生成
// 在左括号和右括号都用完时候递归终止
//在左括号大于右括号时候递归终止
//不需要回溯
var rs []string

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	rs = []string{}

	backTrackgenerateParenthesis("", n, n)
	return rs

}

/**
回溯递归
参数

*/
func backTrackgenerateParenthesis(digits string, left, right int) {

	//递归终止条件，存储结果
	if left == 0 && right == 0 {
		rs = append(rs, digits)
		return
	}
	if left > right {
		return
	}

	//递归逻辑

	if left > 0 {
		backTrackgenerateParenthesis(digits+"(", left-1, right)
	}
	if right > 0 {
		backTrackgenerateParenthesis(digits+")", left, right-1)
	}

}

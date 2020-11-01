package suanfa

import "golang.org/x/tools/go/ssa/interp/testdata/src/strings"

/**
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True

小巧取胜，假设
s2是s1旋转的字符串，长度是一样的
在某个点位变成两个子字符串
它只是将分割点之前的字符串放到后面，自连的话就能又成为头部了
例如
erbottle wat
如果两个s2连接能找到s1的影子（wat放后面，自连又重新连上erbottle）
erbottlewat  erbottlewat


get到没
*/

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	return strings.Contains(s2+s2, s1)
}

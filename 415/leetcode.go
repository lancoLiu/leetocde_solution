package main

import "strconv"

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

/**
num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
*/
func addStrings(num1 string, num2 string) string {
	//字符串形式的非负整数，没有负号，没有前导0
	//从后往前加，使用双指针形式记录位置
	lens1 := len(num1) - 1
	lens2 := len(num2) - 1
	//结果
	ans := ""
	//进位
	add := 0
	//保证只要i,j或者add有值就继续相加
	for i, j := lens1, lens2; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			//巧妙计算出转换值
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}

	return ans
}

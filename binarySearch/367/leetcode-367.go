package main

//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//https://leetcode-cn.com/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	l, r := 2, num/2
	var tmp int
	var mid int
	for l <= r {
		mid = (r-l)/2 + l
		tmp = mid * mid
		if tmp > num {
			r = mid - 1
		} else if tmp == num {
			return true
		} else {
			l = mid + 1
		}
	}
	return false

}

//牛顿法
//x^2 - num = 0 是一个可导函数
//假设取Xk为根的近似值，在(Xk,f(Xk)做切线与x轴相交经过Xk+1)
//Xk + 1 = 1/2(Xk + num / Xk)
func isPerfectSquare2(num int) bool {
	if num < 2 {
		return true
	}

	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}
	return (x*x == num)
}

func main() {

}

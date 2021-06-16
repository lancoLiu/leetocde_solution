package main

//你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币
//给定一个数字 n，找出可形成完整阶梯行的总行数。
//n 是一个非负整数，并且在32位有符号整型的范围内。

//n枚硬币排列的行数不可能超过n行

func arrangeCoins(n int) int {
	//1 1
	//2 3
	//3 6
	//4 10
	//5 15
	//(1+n)*x = 2n
	lo, hight := 0, n
	for lo <= hight {
		mid := lo + (hight-lo)/2
		cost := (mid + 1) * mid / 2
		if cost == n {
			return mid
		} else if cost < n {
			//排mid行的硬币数没有超过n枚，还可以继续排
			lo = mid + 1
		} else {
			hight = mid - 1
		}
	}
	return hight
}

func main() {

}

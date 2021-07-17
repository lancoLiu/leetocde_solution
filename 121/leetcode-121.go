package main

//dp数组
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		//今天为止不持股，可能是卖了，或者本来就没买
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		//今天为止持有股票
		dp[i][1] = max(-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

package main

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//dp动态规划
//dp方程f(i) = max(f(i - 1)+nums[i],nums[i])
//https://leetcode-cn.com/problems/maximum-subarray/solution/dong-tai-gui-hua-fen-zhi-fa-python-dai-ma-java-dai/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//当前最大的子序和
	maxNum := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		maxNum = max(maxNum, nums[i])

	}
	return maxNum
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

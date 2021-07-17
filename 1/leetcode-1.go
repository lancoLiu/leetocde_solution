package main

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//返回数组的位置
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	maps := make(map[int]int)
	for k, v := range nums {
		a := target - v
		if kk, ok := maps[a]; ok {
			return []int{k, kk}
		} else {
			maps[v] = k
		}
	}
	return nil
}

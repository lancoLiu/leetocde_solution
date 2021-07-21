package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l-r)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//mid 大于target,移动右边区间,因为搜索是l<=r，[l,r]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

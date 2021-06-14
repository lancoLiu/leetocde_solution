package main

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//nums = [5,7,7,8,8,10], target = 8
//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/si-lu-hen-jian-dan-xi-jie-fei-mo-gui-de-er-fen-cha/
func searchRange(nums []int, target int) []int {
	var res []int
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	temp := binarySearchLeftBound(nums, target)
	if temp == -1 {
		return []int{-1, -1}
	}
	res = append(res, temp)
	temp = binarySearchRightBound(nums, target)
	res = append(res, temp)
	return res
}

func binarySearchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//[right+1,right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//边界判断
	if left != len(nums) && nums[left] == target {
		return left
	}
	return -1
}

/**
搜查右侧边界
*/
func binarySearchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

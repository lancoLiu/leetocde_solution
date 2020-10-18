package BinarySearch

//https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return left
	} else if right < 0 || nums[right] != target {
		return 0
	}
	return 0
}

package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l >= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		//
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l

}
func main() {

}

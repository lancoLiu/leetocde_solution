package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

func search(nums []int, target int) int {
	l, mid, r := 0, 0, len(nums)-1

	for l <= r {
		mid = (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	//[r,l]
	return -1
}

func main() {

}

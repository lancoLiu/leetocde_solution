package BinarySearch

//https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

/**
!有序是前提
返回数组下标
0-2-4-6-8-7-5-3-1
返回1
*/
//只有中位数小于左边，递减
// 中位数小于右边，递增
func peakIndexInMountainArray(arr []int) int {
	right := len(arr) - 1
	left := 0
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

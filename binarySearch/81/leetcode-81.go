package main

//搜索旋转排序数组，有重复值
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] {
			l++
			continue
		}

		if nums[mid] >= nums[l] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1

			}

		} else {
			//后半段有序
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1

			} else {
				r = mid - 1

			}

		}
	}
	return false
}
func main() {

}

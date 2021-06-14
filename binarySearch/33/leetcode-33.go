package main

//整数数组 nums 按升序排列，数组中的值 互不相同 。

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		//[l,mid]是有序的,必须用等于在[3,1]；target为1情况下，mid和l下标一样，就进不去第一个逻辑
		if nums[mid] >= nums[l] {
			//target在[l,mid)范围内
			if target >= nums[l] && target < nums[mid] {
				//因为for条件是[l,r]，所以根据上述分析target的范围，下一个寻找范围为[l,mid-1]
				r = mid - 1
			} else {
				//否则target在[mid+1，r]中的范围去寻找
				l = mid + 1
			}
		} else {
			//nums[mid] < nums[l]。右边的是有序数组
			//target的范围为(mid,r]
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}

func main() {

}

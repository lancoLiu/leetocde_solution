package BinarySearch

//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/da-jia-bu-yao-kan-labuladong-de-jie-fa-fei-chang-2/
//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/si-lu-hen-jian-dan-xi-jie-fei-mo-gui-de-er-fen-cha/
func SearchRange(nums []int, target int) []int {
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

//左闭右开
//左闭有闭

//while(left < right) 这种写法表示在循环体内部排除元素；
//退出循环的时候 left 和 right 重合，区间 [left, right] 只剩下成 11 个元素，这个元素 有可能 就是我们要找的元素。

//while(left <= right) 这种写法表示在循环体内部直接查找元素；
//退出循环的时候 left 和 right 不重合，区间 [left, right] 是空区间。

/**
搜查左侧边界
第二次二分查找是right不断左移以确定target的最后一个位置，如果没有符合target的值，最后搜索到left=0,right=1时的情况：这时mid=0，nums[mid]!=target，则right=mid-1=-1,溢出。所以加1就能避免溢出。第一次二分查找自己仔细分析一下，并不会出现溢出。
*/

/**
搜查左侧边界
*/
func binarySearchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//[right+1,right]
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid
		}
	}
	//边界判断
	if nums[left] != target {
		return -1
	}
	return left
}

func binarySearchLeftBound2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//[right+1,right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			//应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		} else if nums[mid] > target {
			// 此时 nums[mid] > target，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		} else {
			// ① 不可以直接返回，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		}
	}
	// 此时 left 和 right 的位置关系是 [right, left]，注意上面的 ①，此时 left 才是第 1 次元素出现的位置
	// 因此还需要特别做一次判断
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
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// 别返回，锁定右侧边界
			left = mid
		}
	}
	return left
}

func binarySearchRightBound2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		} else if nums[mid] > target {
			// 此时 nums[mid] > target，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		} else {
			// 只有这里不一样：不可以直接返回，应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		}
	}
	// 由于 findFirstPosition 方法可以返回是否找到，这里无需单独再做判断
	return right
}

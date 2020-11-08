package hot100

import "math"

/**
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, n] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:

输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]

第一次遍历在值对应的位置上数字设置为当前数字乘-1
第二次在遍历数组，如果值不是负数说明这个索引代表的数字不存在
*/
func findDisappearedNumbers(nums []int) []int {

	var res = []int{}

	for _, value := range nums {
		index := int(math.Abs(float64(value)) - 1)
		if nums[index] > 0 {
			nums[index] = nums[index] * -1
		}

	}

	for key, value := range nums {
		if value > 0 {
			res = append(res, key+1)
		}
	}

	return res

}

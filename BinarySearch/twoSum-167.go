package BinarySearch

//https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
/**
输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
*/
//二分查找
func twoSum(numbers []int, target int) []int {
	//边界处理
	var res []int

	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			res = append(res, []int{left + 1, right + 1}...)
			return res
		}
	}
	return res
}

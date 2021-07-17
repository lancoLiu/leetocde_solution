package main

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

//注意：答案中不可以包含重复的三元组。

//不重复意味着第二重循环枚举元素不小于第一重
//第三重循环枚举元素不小于第二重

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]

		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同，second > first+1为必须，否则[-1,-1,2]就不行
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//保证b在c的左侧
			for second < third && nums[second]+nums[third] > target {
				//数太大了
				third--
			}
			//指针重合
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

//

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r && l < len(nums)-1 {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l += 1
			} else if s > 0 {
				r -= 1
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			}
		}
	}
	return res
}

package Backtracking

import "sort"

/**
https://leetcode-cn.com/problems/permutations-ii/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liwe-2/
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
var res3 [][]int

func permuteUnique(nums []int) [][]int {
	var trackback []int
	res3 = [][]int{}
	var used = make([]bool, len(nums))
	sort.Ints(nums)
	dfs2(nums, trackback, used)
	return res3

}

/**

 */
func dfs2(nums, trackback []int, used []bool) {
	if len(trackback) == len(nums) {
		res3 = append(res3, trackback)
		return
	}

	for i := 0; i < len(nums); i++ {
		//这里需要注意的是i-1为false意思是i-1刚刚被撤销了，这里剪枝
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		used[i] = true
		// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
		trackback = append(trackback, nums[i])
		dfs(nums, trackback, used)
		// 撤销刚才的选择，也就是恢复操作
		trackback = trackback[:len(trackback)-1]
		// 标记成未使用
		used[i] = false
	}

}

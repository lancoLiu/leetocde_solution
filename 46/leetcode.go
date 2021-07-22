package main

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	if len(nums) == 0 {
		return nil
	}
	trackback := []int{}
	used := make([]bool, len(nums))

	dfs(nums, trackback, used)
	return res

}

func dfs(nums []int, trackback []int, used []bool) {
	if len(trackback) == len(nums) {
		//添加
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, trackback)
		// 把本次结果追加到最终结果上
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
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
}

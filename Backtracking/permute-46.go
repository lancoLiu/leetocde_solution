package Backtracking

import "fmt"

var res2 [][]int

func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var trackback []int
	var used = make([]bool, len(nums))
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	res2 = [][]int{}
	dfs(nums, trackback, used)
	return res2
}

func dfs(nums []int, trackback []int, used []bool) {
	if len(trackback) == len(nums) {
		res2 = append(res2, trackback)
		return

	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			trackback = append(trackback, nums[i])
			fmt.Printf("i的值:%v,添加了的值:%v,切片的值:%v\n", i, nums[i], trackback)
			dfs(nums, trackback, used)
			// 撤销刚才的选择，也就是恢复操作
			fmt.Printf("回退操作，i的值:%v,删除了值:%v,切片的值:%v\n", i, nums[i], trackback)
			trackback = trackback[:len(trackback)-1]
			// 标记成未使用
			used[i] = false
		}

	}
}

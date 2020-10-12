package Backtracking

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/binary-watch/

func readBinaryWatch(num int) []string {
	result := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			b1 := fmt.Sprintf("%b", i)
			b2 := fmt.Sprintf("%b", j)
			sumOne := strings.Count(b1, "1") + strings.Count(b2, "1")
			if sumOne == num {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}

		}
	}
	return result
}

func readBinaryWatch2(num int) []string {
	res := []string{}
	// []int{1, 2, 4, 8, 16, 32, 1, 2, 4, 8} // 0-5 min 6-9 h 10盏灯
	backTrack(0, 10, num, []int{}, &res)
	return res
}

//backTrack 回溯算法 start 循环开始，从上一个递归+1开始以免找到重复的, cap 灯的总数量, target 目标亮灯的数量 path []int  亮灯的index集合(记录回溯的路径) res 指针返回结果
func backTrack(start, cap, target int, path []int, res *[]string) {
	if len(path) == target { //选择亮灯的数量达到target
		min, hour := 0, 0
		for _, v := range path {
			if v >= 4 { // 如果灯的下表超过3 则表示这些灯是 minute的灯
				min += 1 << (v - 4)
			} else { //如果 灯的index 小于4 表示是 hour的灯亮
				hour += 1 << v
			}
		}
		if min < 60 && hour < 12 { //排除不符合规则的,判断min hour 值是否有效
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, min)) //格式化拼接成字符串
		}
		path = []int{} //重置 回溯的path
		return
	}
	for i := start; i < cap; i++ {
		path = append(path, i)
		backTrack(i+1, cap, target, path, res) //注意这里是 i+1
		path = path[:len(path)-1]              //对 path 进行rollback
	}
}

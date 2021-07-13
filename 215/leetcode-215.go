package main

import (
	"math/rand"
	"time"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5

//快排思想
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	//从小到大排序
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	} else {
		return quickSelect(a, l, q-1, index)
	}
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	//随机化，避免因为顺序或者倒序数组带来的算法退化问题
	//随机i与r交换

	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
	//return
}

//开始快速排序
func partition(a []int, l, r int) int {

	i := l - 1
	//循环，将大于a[r]的放在右边，小于a[r]的在它左边
	//目前是在i左边都是小于a[r]的，所以在最后需要将a[r]调到a[i+1]中并返回i+1
	for j := l; j < r; j++ {

		if a[j] <= a[r] {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

//-----------------------------------------------------

//优先队列思想
//func findKthLargest2(nums []int, k int) int {

//}

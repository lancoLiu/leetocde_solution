package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(a []int, l, r int) {
	if l < r {
		pos := randomPartition(a, l, r)
		quickSort(a, l, pos-1)
		quickSort(a, pos+1, r)
	}
}
func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}
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

func main() {
	a := []int{5, 2, 3, 1}
	aa := sortArray(a)
	fmt.Printf("%#v", aa)
}

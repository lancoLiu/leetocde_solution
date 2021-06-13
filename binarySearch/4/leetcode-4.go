package main

import "fmt"

//暴力法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//暴力解法，合并两个数组
	length1 := len(nums1)
	length2 := len(nums2)

	if length1 == 0 && length2 == 0 {
		return 0
	}
	if length1 == 0 {
		if length2%2 == 0 {
			return float64(nums2[length2/2-1]+nums2[length2/2]) / 2.0
		} else {
			return float64(nums2[length2/2])
		}
	}

	if length2 == 0 {
		if length1%2 == 0 {
			return float64(nums1[length1/2-1]+nums1[length1/2]) / 2.0
		} else {
			return float64(nums1[length1/2])
		}
	}

	//开始合并两个有序数组

	start1 := 0
	start2 := 0
	var res []int
	for len(res) != length1+length2 {
		if start2 >= length2 {
			for start1 < length1 {
				res = append(res, nums1[start1])
				start1++
			}
			break
		}

		if start1 >= length1 {
			for start2 < length2 {
				res = append(res, nums2[start2])
				start2++
			}
			break
		}

		if nums1[start1] > nums2[start2] {
			res = append(res, nums2[start2])
			start2++
		} else {
			res = append(res, nums1[start1])
			start1++
		}
	}

	if len(res)%2 == 0 {
		return float64((res[len(res)/2-1] + res[len(res)/2])) / 2.0
	} else {
		return float64(res[len(res)/2])
	}

}

//二分查找

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	//中位数是两个有序数组中的第 (m+n)/2 个元素
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		//当 m+nm+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2个元素和第 (m+n)/2+1 个元素的平均值。
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

/**
根据二分查找第k个小的数
查找数组A的第k / 2小与B数组的第k / 2小
1、如果A[k/2 - 1] < B[k /2 -1]
说明A数组的前k/2-1数都不是答案，可以排除掉
排除之后两个数组的大小变了，k值更新，因为排除掉A数组的k/2个元素。所以新的k应该为k-k/2

2、如果B[k/2 - 1] < A[k / 2 - 1],如上分析

3、A[k/2 - 1] == B[k /2 -1]，归入第一种情况

考虑上面三种情况的特殊情况
1、A[k/2-1]或者B[k/2-1]越界了，那就选择对应数组的最后一个元素，再根据排除数做减法
//例如oldIndex1 + k/2是越界的，那我取A数组的最后元素index1 = len(A)-1,
如果A[index1] <= B[index2]（index2没有越界）,那相应的k应该更新为k = k - (index1-oldIndex1 + 1)

2、第二种如果一个数组为空了，即index1 = len(A)，就可以返回另一个数组的第k小元素
*/
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			//考虑A[k / 2 - 1]越界了，根据排除数减少个数，而不能直接将K减去k/2
			k -= (newIndex1 - index1 + 1)
			//往右移动，可能超出数组大小
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(a1, a2))
}

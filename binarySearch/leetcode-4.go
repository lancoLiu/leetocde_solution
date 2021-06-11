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

func main() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(a1, a2))
}

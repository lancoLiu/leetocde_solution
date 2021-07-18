package main

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

//从倒序开始比较，nums最后一个元素为m-1，nums2最后一个元素为n-1
func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		//如果nums1比nums2大，意思是nums1的元素要移到最后
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			//
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	//这里只需要考虑到m为0，意思nums1已经全部已经按顺序排好了，就差nums2了，nums2可能也已经排好了
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}

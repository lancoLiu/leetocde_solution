package main

//有序数组中至少h篇论文被引用h次
//找到一个数，使其数值等于右边的数个数
func hIndex(citations []int) int {
	length := len(citations)

	l, m, r := 0, 0, length-1

	for l <= r {
		m = (r-l)/2 + l

		if citations[m] == length-m {
			return length - m
		} else if citations[m] < length-m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return length - l

}

func main() {

}

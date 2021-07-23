package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	//左边
	l := 0
	//右边界
	r := len(matrix[0]) - 1
	//上边界
	t := 0
	//下边界
	b := len(matrix) - 1

	//
	var res []int

	for {
		//从左到右打印
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		//从上到下打印
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		//从右到左打印
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if b < t {
			break
		}
		//下到上
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return res

}

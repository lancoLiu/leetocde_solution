package Tree

//左右中
//中左右
//先来个倒序存储变为右左中
//再调换入栈位置变为左右中
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{curr.Val}, res...)
		for i := 0; i < len(curr.Children); i++ {
			stack = append(stack, curr.Children[i])
		}

	}
	return res
}

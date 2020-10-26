package Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归解法
//https://leetcode-cn.com/problems/symmetric-tree/solution/hua-jie-suan-fa-101-dui-cheng-er-cha-shu-by-guanpe/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsSymmetric(root.Left, root.Right)
}

func dfsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfsSymmetric(left.Left, right.Right) && dfsSymmetric(left.Right, right.Left)

}

//迭代解法
/**
队列保存，一开始将root进队列两次
之后每次取两个作对比
依次将left的left和right.right进队列
left.right和right.left
*/
func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)

	for len(q) > 0 {
		u, v := q[0], q[1]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)

	}
	return true
}

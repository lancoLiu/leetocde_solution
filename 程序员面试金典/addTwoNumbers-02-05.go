package suanfa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。



示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912

*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var dummy = &ListNode{}
	res := dummy
	for l1 != nil || l2 != nil {
		sum := 0
		sum += carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {

			sum += l2.Val
			l2 = l2.Next
		}

		res.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		res = res.Next
		carry = sum / 10

	}
	if carry != 0 {
		res.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummy.Next
}

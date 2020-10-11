package LinkedList

//https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/submissions/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//链表本身数字是高位到低位，所以可以在遍历同时计算
func getDecimalValue(head *ListNode) int {
	curr := head
	var ans int
	for curr != nil {
		ans = ans*2 + curr.Val
		curr = curr.Next
	}
	return ans
}

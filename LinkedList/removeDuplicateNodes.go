package LinkedList

//https://leetcode-cn.com/problems/remove-duplicate-node-lcci/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	var pre, cur *ListNode
	var v bool
	flag := make(map[int]bool)
	for cur = head; cur != nil; {
		v = flag[cur.Val]
		if v {
			pre.Next, cur = cur.Next, cur.Next
			continue
		}
		flag[cur.Val] = true
		pre, cur = cur, cur.Next
	}
	return head
}

//时间换空间，不使用额外空间
func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head
}

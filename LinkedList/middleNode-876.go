package LinkedList

//https://leetcode-cn.com/problems/middle-of-the-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//快慢指针

func MiddleNode(head *ListNode) *ListNode {
	var fast, slow = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

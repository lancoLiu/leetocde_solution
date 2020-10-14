package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//递归版本
func reverseListDFS(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverseListDFS(head.Next)
	head.Next.Next = head
	//防止链表循环，需要将head.next设置为空
	head.Next = nil
	//每层递归函数都返回cur，也就是最后一个节点
	return curr

}

//迭代版本
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre

}

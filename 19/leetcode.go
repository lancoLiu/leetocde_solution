package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

//进阶：你能尝试使用一趟扫描实现吗？
//[1,2] => 1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head

	slow, fast := dummy, head

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

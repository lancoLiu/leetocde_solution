package LinkedList

func ReverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}

	/**
	思路：
	当俩俩一组反转时候
	例如1-2-3-4
	1、cur代表2，head代表1
	2、先存储下一个递归反转的头节点
	3、将当前的两个节点反转返回2-1
	4、反转后连接k组反转的链表，总之head就是指向第一个节点，所以head.Next
	5、返回cur是因为经过俩节点反转后cur在前头，head在后头
	*/
	next := cur.Next
	cur.Next = nil
	cur = reverse2(head)
	head.Next = ReverseKGroup(next, k)

	return cur

}

//旋转链表的写法
func reverse2(head *ListNode) *ListNode {
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

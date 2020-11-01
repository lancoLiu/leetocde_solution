package suanfa

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
两个指针，ob指向head，oc遍历一次链表去除与ob相同的值节点
ob移动一个，oc继续遍历
*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ob := head
	for ob != nil {
		oc := ob
		for oc.Next != nil {
			if ob.Val == oc.Next.Val {
				oc.Next = oc.Next.Next
			} else {
				oc = oc.Next
			}
		}
		ob = ob.Next
	}
	return head
}

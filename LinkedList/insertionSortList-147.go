package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//插入排序

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	//tail指向排好序的尾节点，sort是下一个待排序的节点
	tail := head
	sort := head.Next

	for sort != nil {
		if sort.Val < tail.Val {
			//每次重新重dummy出发找到本节点该插入的位置
			pos := dummy
			for pos.Next.Val < sort.Val {
				pos = pos.Next
			}
			//dummy-》4-》1-》3-》2为例子

			//排好序的尾节点指向待排序节点的下一个节点（4指向1，2为本轮待排序节点）
			tail.Next = sort.Next
			//pos此时在dummy处，先让这轮排序的节点连接它的下一个节点，以免丢失。即2连接4
			sort.Next = pos.Next
			//dummy-》2
			pos.Next = sort
			//下一个待排序节点
			sort = tail.Next
		} else {
			tail = tail.Next
			sort = sort.Next
		}
	}
	return dummy.Next
}

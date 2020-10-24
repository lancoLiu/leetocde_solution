package LinkedList

//重排链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
1、找到中间节点
--这里的中间节点如果是偶数就找后面那个才对
2、反转后半段
3、前半断与后半断依次拼接
*/
func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	//1,2
	mid := midNode(head)
	l1 := head

	//反转之后，后半段是
	l2 := reverse(mid)
	mid.Next = nil

	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next
		l1.Next = l2
		l1 = l1Tmp
		l2.Next = l1
		l2 = l2Tmp
	}

}

func midNode(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {

	//链表反转
	pre := head
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre

}

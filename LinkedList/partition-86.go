package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。



示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/
//遍历原链表，用了一个队列存储所有大于x的元素
//如果小于x则直接插入新链表中，直到遍历完成在遍历队列拼接元素即可
func Partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	var q []int
	for head != nil {
		if head.Val < x {
			curr.Next = head
			curr = curr.Next
		} else {
			q = append(q, head.Val)
		}
		head = head.Next
	}
	for _, value := range q {
		curr.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		curr = curr.Next
	}
	return dummy.Next

}

//维持两个链表两个指针
func partition2(head *ListNode, x int) *ListNode {
	before_head, after_head := &ListNode{}, &ListNode{}
	before, after := before_head, after_head

	for head != nil {

		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	//实现断环
	//假设【1，3，5，6，2】x=3，拿到after是【3，5，6】但是6next还是指向2的
	after.Next = nil
	before.Next = after_head.Next
	return before_head.Next

}

package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

*/
//采用分治或者归并排序

func SortList(head *ListNode) *ListNode {

	length := getLength(head)
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummy.Next = head
	for step := 1; step < length; step *= 2 {
		pre := dummy
		cur := dummy.Next
		for cur != nil {
			h1 := cur
			h2 := cut(h1, step)
			cur = cut(h2, step)
			temp := merge3(h1, h2)
			pre.Next = temp
			for pre.Next != nil {
				pre = pre.Next
			}

		}

	}
	return dummy.Next

}

//俩链表有序合并
func merge3(l1, l2 *ListNode) *ListNode {
	rs := &ListNode{}
	cur := rs

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l2 != nil {
		cur.Next = l2
	}
	if l1 != nil {
		cur.Next = l1
	}
	return rs.Next
}

//切分点
func cut(head *ListNode, step int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for i := 1; i < step && cur.Next != nil; i++ {
		cur = cur.Next
	}
	right := cur.Next
	cur.Next = nil
	return right
}

//求链表长度
func getLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	return count
}

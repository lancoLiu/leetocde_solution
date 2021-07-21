package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//简单分析：
/**
当链表无环，快指针会走到nil，直接返回nil即可
假设下面情况都是有环：
当快慢指针第一次相遇的时候，因为快指针f走的路程是慢指针s的2倍
即f = 2s
假设环长度为b，直线长度为a
f比s多走了n个环长度(这里n不确定)
f = s + nb
即2s(f) = s +nb
所以第一次相遇时候：
s = nb
f = 2nb

又假设从链表头节点走到链表环入口时候是a+nb（n等于0刚好没绕环）
所以在第一次相遇时候让f重新从头节点开始即可
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var intersect *ListNode = getFirstMeet(head)
	if intersect == nil {
		return nil
	}
	var ptr1 *ListNode = head
	var ptr2 *ListNode = intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

func getFirstMeet(head *ListNode) *ListNode {
	f, s := head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			return s
		}
	}

	return nil
}

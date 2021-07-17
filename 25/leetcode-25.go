package main

//K个一组反转链表

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	//先划分出那一组
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	//如果凑不足一组，直接返回了
	if cur == nil {
		return head
	}
	//保存当前组的后驱节点
	next := cur.Next
	//断链，这样可以将当前组当成反转链表题去写
	cur.Next = nil
	//传head,
	//1-2-3-4,传递1过去，返回时候是返回2-1，即cur为2这个节点
	cur = reverse(head)
	//为什么这里能递归调用，因为除开当前组的链表是个子问题，它也是可以继续k个一组断链再反转的。
	//继续递归调用next这条链表，将返回值链接到head的next上即1之后
	head.Next = reverseKGroup(next, k)
	//返回cur也容易懂，返回反转之后的头节点
	return cur
}

//迭代法旋转链表
//1->2->3
func reverse(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	pre := &ListNode{}
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre

}

func main() {

}

//迭代法
//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	//dummy
	dummy := &ListNode{Next: head}
	//pre 代表待翻转链表的前驱，end 代表待翻转链表的末尾
	pre := dummy
	end := dummy

	for end.Next != nil {
		//i从0开始，因为end是dummy节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start := pre.Next
		//经过k此循环，end 到达末尾，记录待翻转链表的后继 next = end.next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}
	return dummy.Next

}

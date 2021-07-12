package main

import "lanco.vip/lib"

//https://leetcode-cn.com/problems/reverse-linked-list/solution/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/
func reverseList(head *lib.ListNode) *lib.ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	//三步
	//
	newHead := reverseList(head.Next)
	//反转，后一个节点的next指向前一个节点
	head.Next.Next = head
	//断链，当前节点的next断掉
	head.Next = nil
	//返回，这里的newHead实际上一直指向最后一个
	return newHead

}

//迭代
func reverseList2(head *lib.ListNode) *lib.ListNode {

	if head == nil {
		return nil
	}
	//当前节点的前一个节点，初始为空
	var pre *lib.ListNode
	cur := head

	for cur != nil {
		//存储当前节点的下一个节点
		tmp := cur.Next
		//当前节点指向前一个节点
		cur.Next = pre
		//前一个节点开始向后移动
		pre = cur
		//当前节点开始向后移动
		cur = tmp
	}
	return pre

}

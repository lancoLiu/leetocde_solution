package LinkedList

//https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4

*/
//快慢指针，快指针先走k步，慢指针再开始走，这样快指针走到末尾，慢指针就是倒数第k个节点了
//2
//1-2-3-4

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	fast, slow := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

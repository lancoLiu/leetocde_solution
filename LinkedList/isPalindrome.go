package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
编写一个函数，检查输入的链表是否是回文的。



示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true

*/

//用 O(n) 时间复杂度和 O(1) 空间
//先找到中间节点
//对后半段节点反转
//对比反转后的和head开始即可

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	//1,2,3,4
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	reverseLink := reverse(slow)

	for reverseLink != nil {
		if reverseLink.Val != head.Val {
			return false
		}
		reverseLink = reverseLink.Next
		head = head.Next
	}
	return true

}

func reverse(node *ListNode) *ListNode {
	var pre *ListNode
	next := &ListNode{}
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

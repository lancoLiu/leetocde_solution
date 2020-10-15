package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//递归求解，三步走
//1、递归返回值
//2、递归终止条件
//3、本层递归要做的逻辑操作

//1-2->3->4
//2->1->4->3
func swapPairs(head *ListNode) *ListNode {

	//递归终止条件
	if head == nil || head.Next == nil {
		//设置递归返回值
		return head
	}
	//递归逻辑操作

	tmp := head.Next
	head.Next = swapPairs(tmp.Next)
	tmp.Next = head
	return tmp
}

//迭代版本
func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	tmp := dummy

	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		//注意这一步
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return dummy.Next
}

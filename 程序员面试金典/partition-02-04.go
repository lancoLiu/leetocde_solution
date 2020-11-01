package suanfa

/**
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

示例:

输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
*/
//双指针技巧，q一直往后移动直到末尾，p是指向第一个>=x的值，以备交换
func Partition(head *ListNode, x int) *ListNode {
	p, q := head, head
	for q != nil {
		if q.Val < x {
			q.Val, p.Val = p.Val, q.Val
			p = p.Next
		}
		q = q.Next
	}
	return head
}

package LinkedList

//https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
node为要删除的节点，我们可以将node与下一个节点交换
然后再跳过下一个节点即可看为删除

*/
func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

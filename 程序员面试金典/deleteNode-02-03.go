package suanfa

/**
删除中间节点
这里是直接替换值了
例如1-2-3-4
要替换3
这里是将4赋值给3这个节点
然后跳过原本4这个节点
*/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

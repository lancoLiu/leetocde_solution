package suanfa

/**
环路检测
使用第二种方法，存在哈希集合里面
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var seen = map[*ListNode]*ListNode{}

	for head != nil {
		if _, ok := seen[head]; ok {
			return seen[head]
		}
		seen[head] = head
		head = head.Next
	}
	return nil
}

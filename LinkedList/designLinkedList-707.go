package LinkedList

//链表,大写,json
type ListNode struct {
	Value int       `json:"val"`
	Next  *ListNode `json:"next"`
}
type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	lens int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.lens {
		return -1
	}
	head := this.head
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head.Value

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	/**
	var head MyLinkedList = MyLinkedList{}
	head.Value = val
	head.Next = this
	this = &head
	**/

	this.head = &ListNode{
		Value: val,
		Next:  this.head,
	}
	if this.lens == 0 {
		this.tail = this.head
	}
	this.lens++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.lens == 0 {
		this.tail = &ListNode{
			Value: val,
			Next:  nil,
		}
		this.head = this.tail
	} else {
		this.tail.Next = &ListNode{
			Value: val,
			Next:  nil,
		}
		this.tail = this.tail.Next
	}
	this.lens++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.lens {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.lens {
		this.AddAtTail(val)
		return
	}
	head := this.head
	for i := 1; i < index; i++ {
		head = head.Next
	}
	head.Next = &ListNode{
		Value: val,
		Next:  head.Next,
	}
	this.lens++

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.lens {
		return
	}
	if index == 0 {
		this.DeleteAtHead()
		return
	}
	if index == this.lens-1 {
		this.DeleteAtTail()
		return
	}
	head := this.head
	for i := 1; i < index; i++ {
		head = head.Next
	}
	head.Next = head.Next.Next
	this.lens--
}

func (this *MyLinkedList) DeleteAtHead() {
	if this.lens == 0 {
		return
	}
	this.head = this.head.Next
	if this.lens == 1 {
		this.tail = this.head
	}
	this.lens--
}

func (this *MyLinkedList) DeleteAtTail() {
	if this.lens == 0 {
		return
	}
	head := this.head
	for head.Next != this.tail {
		head = head.Next
	}
	this.tail = head
	if this.lens == 1 {
		this.head = this.tail
	}
	this.lens--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

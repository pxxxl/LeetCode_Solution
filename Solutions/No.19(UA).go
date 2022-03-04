package main

//hash 版本
func removeNthFromEndA(head *ListNode, n int) *ListNode{
	dir := make([]*ListNode, 1)
	dir[0] = head
	tra := head
	for tra.Next != nil{
		dir = append(dir, tra.Next)
		tra = tra.Next
	}
	num := len(dir) - n
	tra = dir[num]
	tra.Next = tra.Next.Next
	return head
}

//双指针版本
func removeNthFromEndB(head *ListNode, n int) *ListNode{
	step := n+1
	forward := head
	backward := head
	for step != 0{
		forward = forward.Next
		step--
	}
	for forward.Next == nil{
		forward = forward.Next
		backward = backward.Next
	}
	backward.Next = backward.Next.Next
	return head
}

package main

//最强的人可以直接在原链表上操作
func addTwoNumbersA(l1 *ListNode, l2 *ListNode) *ListNode {
	traL := l1
	traS := l2
	cur := 0
	flag := 0
	run := 0
	head := ListNode{Next: nil, Val: 0}
	tra := &head
	first := 0
total:
	for {
		switch {
		case traL == nil && traS == nil:
			run = 2
		case traL != nil && traS != nil:
			run = 0
		default:
			run = 1
			if traL == nil {
				traL = traS
				traS = nil
			}
		}

		switch run {
		case 0:
			cur = traL.Val + traS.Val + flag
			flag = cur / 10
			cur = cur % 10
			traL = traL.Next
			traS = traS.Next
		case 1:
			cur = traL.Val + flag
			flag = cur / 10
			cur = cur % 10
			traL = traL.Next
		case 2:
			if flag == 1 {
				newNode := ListNode{Next: nil, Val: flag}
				tra.Next = &newNode
				tra = tra.Next
			}
			break total
		}
		if first == 0 {
			head.Val = cur
			first = 1
		} else {
			newNode := ListNode{Next: nil, Val: cur}
			tra.Next = &newNode
			tra = tra.Next
		}
	}
	return &head

}

//原链表上的操作
func addTwoNumbersB(l1 *ListNode, l2 *ListNode) *ListNode{
return nil
}
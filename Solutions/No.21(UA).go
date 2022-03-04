package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := list1
	copy := list1
	first := true
	small := list1
	big := list2
	for list1 != nil&&list2 != nil{
		if list1.Val > list2.Val{
			big = list1
			small = list2
		}else{
			big = list2
			small = list1
		}
		if !first{
			copy.Next = small
		}else{
			head = small
			first = false
		}
		copy = small
		small = small.Next
		copy.Next = big
	}
	return head
}

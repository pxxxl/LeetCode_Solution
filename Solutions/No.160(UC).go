package main

//UC:双指针

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dicA := map[*ListNode]struct{}{}
	for headA != nil{
		dicA[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil{
		_, ok := dicA[headB]
		if ok{
			return headB
		}
		headB = headB.Next
	}
	return nil
}

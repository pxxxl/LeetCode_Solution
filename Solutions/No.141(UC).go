package main

//UC:快慢指针

//散列表
func hasCycle(head *ListNode) bool {
	dir := make(map[*ListNode]struct{})
	for{
		if head.Next == nil{
			return false
		}else{
			_, ok := dir[head.Next]
			if ok{
				return true
			}else{
				dir[head.Next] = struct{}{}
			}
		}
		head = head.Next
	}
}

//快慢指针
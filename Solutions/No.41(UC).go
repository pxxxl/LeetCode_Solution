package main

//UC:交换重排

//Hash
func firstMissingPositiveA(nums []int) int {
	dir := map[int]struct{}{}
	for _, v := range nums{
		_, ok := dir[v]
		if !ok{
			dir[v] = struct{}{}
		}
	}
	i := 1
	for{
		_, ok := dir[i]
		if ok{
			i++
		}else{
			break
		}
	}
	return i
}




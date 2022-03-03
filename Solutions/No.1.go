package main

//高耗时低内存
func twoSum1(nums []int, target int) []int {
	result := make([]int, 2)
	for i, v1 := range nums{
		point := target - v1
		for j := i + 1 ; j < len(nums) ; j++{
			if nums[j] == point {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}

//高内存低耗时
func twoSum2(nums []int, target int) []int {
	result := make([]int, 2)
	hashmap := make(map[int]int)
	for i, v1 := range nums{
		hashmap[target - v1] = i
	}
	for i, v1 := range nums{
		index, exist := hashmap[v1]
		if exist && index != i{
			result[0] = index
			result[1] = i
		}
	}
	return result
}

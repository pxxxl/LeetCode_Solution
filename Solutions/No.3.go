package main

//散列+滑动窗口
func lengthOfLongestSubstring(s string) int {
	head := 0
	tail := 0
	maxNum := 0
	curNum := 0
	for head < len(s){
		if head == tail || s[head] != s[tail] {
			head++
			curNum++
		}else{
			tail++
			curNum--
		}
		if curNum > maxNum {
			maxNum = curNum
		}
	}
	return maxNum
}
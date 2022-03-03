package main

func findSubstring(s string, words []string) []int {
	subStrLen := len(words[0])
	ch := make(chan int, subStrLen)
	for i := 0; i < subStrLen; i++{
		goStr := string(s[i:])
		dir := makeMap(words)
		totalCheck := len(words)
		go check(goStr,dir,subStrLen,totalCheck,i,ch)
	}
	goNum := subStrLen
	result := [1024]int{}
	index := 0
re:
	for{
		select{
		case msg := <- ch:
			if msg == -1{
				goNum--
				if goNum == 0{
					break re
				}
			}else{
				result[index] = msg
				index++
			}
		}
	}
	return result[:index]
}

func makeMap(words []string) (dir map[string]int){
	dir = make(map[string]int)
	for _, v := range words{
		dir[v]++
	}
	return
}

func check(s string, dir map[string]int, subLen int, totalCheck int, bios int, ch chan int){
	head := 0
	tail := 0
	packLen := len(s)/subLen
	for head < packLen {
		headStr := string(s[subLen*head:subLen*head+subLen])
		tailStr := string(s[subLen*tail:subLen*tail+subLen])
		v, ok := dir[headStr]

		switch{
		case head == tail && (!ok || v == 0):
			head++
			tail++
		case head > tail && (!ok || v == 0):
			dir[tailStr]++
			totalCheck++
			tail++
		case ok && v > 0:
			dir[headStr]--
			totalCheck--
			head++
			if totalCheck == 0 {
				ch <- tail*subLen+bios
			}
		}
	}
	ch <- -1
	return
}


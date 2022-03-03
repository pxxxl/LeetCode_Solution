package main

func minWindow(s string, t string) string {
	dir := map[byte]int{}
	sum := 0
	tLen := len(t)
	sLen := len(s)
	H := 0
	T := 0
	for i := 0;i < tLen;i++{
		dir[t[i]]++
		sum++
	}
	re:
	for H < sLen{
		v, ok := dir[s[H]]
		switch{
		case ok && v > 0:
			dir[s[H]]--
			sum--
			H++
			if sum == 0{
				break re
			}

		case H == T && (!ok || v <= 0):
			H++
			T++
		case H > T && (!ok || v <= 0):
			dir[s[T]]++
			sum++
			T++
		}
	}
	return s[T:H-T+1]
}

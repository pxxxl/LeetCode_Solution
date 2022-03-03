package main

func romanToInt(s string) int {
	var flag byte
	sum := 0
	dic := make(map[byte]int)
	for _, v := range s{
		switch v{
		case 'I':
			dic['I'] += 1
			flag = 'I'
		case 'V':
			if flag == 'I'{
				sum -= 2
			}
			dic['V'] += 1
			flag = 'V'
		case 'X':
			if flag == 'I'{
				sum -= 2
			}
			dic['X'] += 1
			flag = 'X'
		case 'L':
			if flag == 'X'{
				sum -= 20
			}
			dic['L'] += 1
			flag = 'L'
		case 'C':
			if flag == 'X'{
				sum -= 20
			}
			dic['C'] += 1
			flag = 'C'
		case 'D':
			if flag == 'C'{
				sum -= 200
			}
			dic['D'] += 1
			flag = 'D'
		case 'M':
			if flag == 'C'{
				sum -= 200
			}
			dic['M'] += 1
			flag = 'M'
		}
	}
	sum += dic['I'] + 5*dic['V'] + 10*dic['X'] + 50*dic['L'] + 100*dic['C'] + 500*dic['D'] + 1000*dic['M']
	return sum
}

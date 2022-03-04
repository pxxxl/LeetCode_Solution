package main

type ListNode struct {
	Val int
	Next *ListNode
}

//规则： #数据@自己序号#数据...   这样轮转
func linkedListBuilder(command string) *ListNode{
	totalData := 0
	numFlag := 0                                     //标一表示正在录入序号
	numBuffer := 0
	addDir := make(map[int]*ListNode)
	for _, v := range command{
		switch v{
		case '#':
			totalData++
			numFlag = 1
		case '@':
			if numFlag == 1{
				var node ListNode
				addDir[numBuffer] = &node
				numFlag = 0
				numBuffer = 0
			}
		}
	}
}
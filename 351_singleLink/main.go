package main

import "fmt"

type heroNode struct {
	no   int
	name string
	next *heroNode
}

func heroAdd(head, newhero *heroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		} else {
			temp = temp.next
		}
	}
	temp.next = newhero
}

// 显示列表
func heroList(head *heroNode) {
	temp := head
	//空链表
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}
	for {

		fmt.Print(temp.next.name, "==>")
		//退出条件
		temp = temp.next
		if temp.next == nil {
			break
		}

	}
}
func main() {
	head := &heroNode{}
	hero1 := &heroNode{
		no:   1,
		name: "宋江",
	}
	hero2 := &heroNode{
		no:   2,
		name: "卢俊义",
	}
	heroAdd(head, hero1)
	heroAdd(head, hero2)

	heroList(head)
}

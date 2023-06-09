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

func heroAdd2(head, newhero *heroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no > newhero.no {
			//说明newHeroNode 就应该插入到 中间
			break
		} else if temp.next.no == newhero.no {
			//说明我们额链表中已经有这个no,就不然插入.
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，id已经存在")
	} else {
		newhero.next = temp.next //如果temp.next.no < newhero.no 这条语句是没用的
		temp.next = newhero
	}
}

func delHero(head, delhero *heroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > delhero.no {
			break
		} else if temp.next.no == delhero.no {
			flag = true
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("删除的节点不存在！")
		return
	}
	temp.next = delhero.next
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
	hero3 := &heroNode{
		no:   3,
		name: "吴用",
	}
	heroAdd2(head, hero1)
	heroAdd2(head, hero3)
	heroAdd2(head, hero2)

	heroList(head)
	fmt.Println()
	fmt.Println("删除节点")
	delHero(head, hero3)
	heroList(head)
	fmt.Println()
	delHero(head, hero3)
}

package main

import "fmt"

type heroNode struct {
	no   int
	name string
	next *heroNode
	pre  *heroNode
}

// 向最后一位添加数据
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
	newhero.pre = temp //往前索引的链
}

// 向中间添加数据
func heroAdd2(head, newhero *heroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no > newhero.no {
			//说明newHeroNode 就应该插入 中间去，temp就不动了，赶紧出来
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
		newhero.next = temp.next //如果temp.next.no < newhero.no 这条语句是没用的  ok
		newhero.pre = temp       //ok

		if temp.next != nil {
			temp.next.pre = newhero
		}
		temp.next = newhero //
	}
}

func delHero(head, delhero *heroNode) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > delhero.no { //这个是多余的，没必要加
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
	//if temp.next != nil {
	//	delhero.next.pre = temp
	//}
	temp.next = delhero.next

}

// 从前往后显示列表
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

// 从后往前显示列表
func heroList2(head *heroNode) {
	temp := head
	//空链表
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}
	//将 temp 放到最后
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		fmt.Print(temp.name, "==>")
		//退出条件
		temp = temp.pre
		if temp.pre == nil {
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
	heroAdd2(head, hero2)
	heroAdd2(head, hero3)

	heroList(head)
	fmt.Println()
	fmt.Println("向前索引打印")
	heroList2(head)
	fmt.Println()
	fmt.Println("删除后的")
	delHero(head, hero3)
	heroList(head)
	fmt.Println()
	fmt.Println("向前索引打印")
	heroList2(head)

}

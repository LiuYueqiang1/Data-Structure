package main

import "fmt"

type hero struct {
	no   int
	name string
	next *hero
}

func addNode(head *hero, newheroNode *hero) {
	tmpe := head
	for {
		if tmpe.next == nil {
			break
		}
		tmpe = tmpe.next
	}
	tmpe.next = newheroNode
}
func readNode(head *hero) {
	if head.next == nil {
		fmt.Println("空")
		return
	}
	tmpe := head
	for {

		fmt.Printf("%d %s ==>", tmpe.next.no, tmpe.next.name)
		tmpe = tmpe.next
		if tmpe.next == nil {
			break
		}
	}
}
func main() {
	head := &hero{}
	hero1 := &hero{
		no:   1,
		name: "宋江",
	}
	hero2 := &hero{
		no:   2,
		name: "卢俊义",
	}
	hero3 := &hero{
		no:   3,
		name: "林冲",
	}
	fmt.Println(head)
	addNode(head, hero1)
	addNode(head, hero2)
	addNode(head, hero3)
	readNode(head)
}

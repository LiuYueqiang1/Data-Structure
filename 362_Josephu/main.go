package main

import (
	"fmt"
)

// 小孩的结构体
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针[默认值是nil]
}

// 编写一个函数，构成单向的环形链表
// num ：表示小孩的个数
// *Boy : 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点

	//判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//分析构成循环链表，需要一个辅助指针[帮忙的]
		//1. 因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环形链表
		}
	}
	return first

}

// 显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {

	//处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	//创建一个指针，帮助遍历.[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		//退出的条件?curBoy.Next == first
		if curBoy.Next == first {
			break
		}
		//curBoy移动到下一个
		curBoy = curBoy.Next
	}
}

/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到 m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/

// 分析思路
// 1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
// 2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {
	fmt.Println()
	//1. 空的链表我们单独的处理
	if first.Next == nil {
		fmt.Println("link is nil!")
		return
	}
	//2、定义辅助指针帮助删除小孩
	tail := first
	//3、给辅助指针形成一个循环，直到tail.next=nil结束
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//4、first移动 startNo位 ，tail移动 startNo位
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		//5、first移动countNum位，tail移动countNum位
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩 %d 出圈\n", first.No)
		//6、first向下移动一位，tail.next指向first
		first = first.Next
		tail.Next = first

		if tail == first {
			fmt.Printf("小孩 %d 出圈\n", first.No)
			break
		}
	}

}

func main() {

	first := AddBoy(5)
	//显示
	ShowBoy(first)
	PlayGame(first, 2, 3)

}

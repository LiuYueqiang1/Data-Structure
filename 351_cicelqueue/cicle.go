package main

import "fmt"

type queue struct {
	head    int
	tail    int
	maxSize int
	array   [5]int
}

// 判断是否为满
func (q *queue) isFull() bool {
	//这样队列才是满的，但是刚开始的时候它俩就相等，加不进去东西，故再加一个判断语句
	if q.size() == q.maxSize {
		return true
	}
	return false
}

// 判断是否为空
func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

// 计算数组中值的个数
func (q *queue) size() int {

	return (q.tail + q.maxSize - q.head) % q.maxSize
}

// add
// get
// show
func main() {
	var val int
	var key string
	que := queue{
		head:    0,
		tail:    0,
		maxSize: 5,
	}
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "add":

		}
	}
}

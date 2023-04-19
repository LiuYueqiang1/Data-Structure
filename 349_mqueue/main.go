package main

import (
	"errors"
	"fmt"
	"os"
)

// 构建队列的结构体
type Queue struct {
	maxSize int
	front   int
	rear    int
	array   [5]int
}

// 添加数据到队列
func (q *Queue) AddQueue(val int) (err error) {
	if q.rear == q.maxSize-1 {
		return errors.New("queue is full")
	}
	q.rear++
	q.array[q.rear] = val
	return
}

// 展示数据队列
func (q *Queue) ShowQueue() (err error) {
	fmt.Println("当前队列为：")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, q.array[i])
	}
	return
}

// 取出一个数
func (q *Queue) DelQueue() (val int, err error) {
	if q.front == q.rear {
		return -1, errors.New("queue empty")
	}
	q.front++ //这个front是结构体中公用的
	val = q.array[q.front]
	return val, err
}
func main() {
	queue := Queue{
		maxSize: 6,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("请输入您的操作:")
		fmt.Println("add，向队列中加值")
		fmt.Println("show，展示队列")
		fmt.Println("delea")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入您要添加的数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("add queue success!")
			}
		case "show":
			queue.ShowQueue()
		case "dele":
			val, err := queue.DelQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("取出的数为：", val)
			}
		case "exit":
			os.Exit(0)

		}
	}
}

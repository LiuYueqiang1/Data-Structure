package main

import (
	"errors"
	"fmt"
)

// 用数组模拟出栈、入栈的过程
type Stack struct {
	stackMax int
	stackTop int //栈顶
	arr      [5]int
}

// 入栈
// 输入一个值，压入栈中。如果超过了栈的最大值,则入栈失败
func (s *Stack) pushStack(val int) (err error) {
	if s.stackTop >= s.stackMax-1 {
		fmt.Println("stack full!")
		return errors.New("stack full!")
	}
	//压栈
	s.stackTop++
	//压入输入的值
	s.arr[s.stackTop] = val
	return
}

// 出栈
// 出栈应该遵循先入后出的原则，从栈顶开始出，如果栈空，则报错
func (s *Stack) popStack() (err error) {
	// 如果栈空了，则显示err，并返回
	if s.stackTop == -1 {
		fmt.Println("stack is nil!")
		return errors.New("stack is nil!")
	}
	// 开始出栈、怎么出栈？
	// 从顶部往外出，故初始值为s.stackTop
	for i := s.stackTop; i >= 0; i-- {
		s.stackTop--
		fmt.Printf("arr[%d]=%d out\n", i, s.arr[i])
	}
	return
}

// 出栈2，一个一个的出
func (s *Stack) popStack2() (val int, err error) {
	// 如果栈空了，则显示err，并返回
	if s.stackTop == -1 {
		fmt.Println("stack is nil!")
		return 0, errors.New("stack is nil!")
	}
	// 开始出栈、怎么出栈？
	// 从顶部往外出，故初始值为s.stackTop
	val = s.arr[s.stackTop]
	fmt.Printf("arr[%d]=%d out\n", s.stackTop, s.arr[s.stackTop])
	s.stackTop--
	return val, nil
}

// 显示栈的列表
func (s *Stack) ListStack() {
	// 如果栈为空的话则显示空，且返回
	fmt.Println("栈中内容为：")
	if s.stackTop == -1 {
		fmt.Println("stack nil!")
		return
	}
	// 遍历这个栈
	//不能这样嗷，不能遍历数组，而是以 s.stackTop 去遍历，所以 以数组遍历是有错误的。
	//for i, v := range s.arr {
	//	fmt.Printf("arr[%d] is %d\n", i, v)
	//}
	//

	//从栈顶开始遍历的
	for i := s.stackTop; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {
	stack := &Stack{
		stackMax: 5,
		stackTop: -1,
	}
	stack.pushStack(1)
	stack.pushStack(2)
	stack.pushStack(6)
	stack.pushStack(7)
	stack.pushStack(10)

	fmt.Println("显示：")
	stack.ListStack()

	fmt.Println("出栈：")
	stack.popStack2()

	fmt.Println("显示：")
	stack.ListStack()

}

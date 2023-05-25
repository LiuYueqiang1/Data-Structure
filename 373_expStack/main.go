package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 用数组模拟出栈、入栈的过程
type Stack struct {
	stackMax int
	stackTop int //栈顶
	arr      [20]int
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
	//	fmt.Printf("arr[%d]=%d out\n", s.stackTop, s.arr[s.stackTop])
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

// 判断一个字符是不是运算符[+,-,*,/]
func (s *Stack) IsOper(val int) bool {

	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算
func (s *Stack) Cal(num1 int, num2 int, oper int) (res int) {
	res = 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		errors.New("运算符无效")
	}
	return res
}

// 判断优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {

	//创建一个数栈、一个符号栈
	numStack := &Stack{
		stackMax: 20,
		stackTop: -1,
	}
	operStack := &Stack{
		stackMax: 20,
		stackTop: -1,
	}

	exp := "30+30*4-3+10"

	num1 := 0
	num2 := 0
	res := 0
	oper := 0 //运算符的ASCII码
	index := 0

	keepNum := ""
	// 如果operStack是一个空栈，则直接入栈
	for {
		// 获取到当前的字母
		ch := exp[index : index+1]
		// 判断是否为运算符
		temp := int([]byte(ch)[0]) // 转换为ASCII码

		//判断是否为符号
		if operStack.IsOper(temp) { //如果是运算符
			if operStack.stackTop == -1 { //如果是空栈
				operStack.pushStack(temp) //直接入栈
			} else { //不是空栈
				// 如果发现opertStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级
				//就从符号栈pop出，并从数栈也 pop 两个数，进行运算，运算后的结果再重新入栈到数栈， 下一个符号入符号栈
				if operStack.Priority(operStack.arr[operStack.stackTop]) >= operStack.Priority(temp) {
					//必须是 >= 如果栈顶的的优先级等于新加入的 temp ，也是先算原来栈中的，否则就变成从栈顶往下算了，先算后面的数了
					//则变成先算 10+3=13 再算120-13=107 ，再算30+107=137
					//所以必须是 >= 保证了先算先入栈的符号，先入栈符号的优先级 >= 后入栈的
					num1, _ = numStack.popStack2()
					num2, _ = numStack.popStack2()
					oper, _ = operStack.popStack2()
					res = operStack.Cal(num1, num2, oper) //得到结果

					//入栈
					numStack.pushStack(res)
					operStack.pushStack(temp) //这个temp指的是下一个符号了
				} else { // 2.2 符号直接入栈
					operStack.pushStack(temp)
				}
			}

		} else { //如果是数
			//处理多位数的思路
			//1.定义一个变量 keepNum string, 做拼接
			keepNum += ch
			//2.每次要向index的后面字符测试一下，看看是不是运算符，然后处理
			//如果已经到表达最后，直接将 keepNum 转化为数字
			if index == len(exp)-1 {
				val, _ := strconv.Atoi(keepNum)
				//压入数字栈
				numStack.pushStack(val)
			} else {
				// 向后找一位看是否为运算符
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					//如果是运算符
					val2, _ := strconv.Atoi(keepNum)
					numStack.pushStack(val2)
					keepNum = ""
				}
			}
		}
		//继续扫描
		//先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	//如果扫描表达式 完毕，依次从符号栈取出符号，然后从数栈取出两个数，
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.stackTop == -1 {
			break
		}
		num1, _ = numStack.popStack2()
		num2, _ = numStack.popStack2()
		oper, _ = operStack.popStack2()
		res = operStack.Cal(num1, num2, oper) //得到结果
		numStack.pushStack(res)
	}
	//如果我们的算法没有问题，表达式也是正确的，则结果就是numStack最后数
	rest, _ := numStack.popStack2()
	fmt.Printf("表达式%s = %v", exp, rest)
}

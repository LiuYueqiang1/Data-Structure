package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//1、写一个原始数组
	var array [11][11]int
	//fmt.Println(array)
	array[1][2] = 1
	array[2][3] = 2
	//1、1打印原始数组
	fmt.Println("原始数组：")
	for _, v1 := range array {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//2、将原始数组改为稀疏数组
	//2、1构建一个结构体存放数据
	type arrNode struct {
		row   int
		col   int
		value int
	}

	var aparray []arrNode
	arrnode := arrNode{
		row:   11,
		col:   11,
		value: 0,
	}
	aparray = append(aparray, arrnode)
	//2、2遍历原始数组，将其值加到结构体中
	for i1, v1 := range array {
		for i2, v2 := range v1 {
			if v2 != 0 {
				arrnode := arrNode{
					row:   i1,
					col:   i2,
					value: v2,
				}
				aparray = append(aparray, arrnode)
			}
		}
	}
	fmt.Println("当前的稀疏数组为：")
	for i, v := range aparray {
		fmt.Printf("%d: %d %d %d \n", i, v.row, v.col, v.value)
		//fmt.Println(i, v.row, v.col, v.value)
	}
	//3、将稀疏数组存到文件中
	//3、1创建一个文件
	filew, err := os.OpenFile("F:\\goland\\go_project\\21weeks\\Data_Structure\\sparsearray\\arr.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0235)
	if err != nil {
		fmt.Printf("open file is failed,err:%v\n", err)
		return
	}
	defer filew.Close()
	//3、2按行
	fw := bufio.NewWriter(filew)
	for i, v := range aparray {
		ss := fmt.Sprintf("%d: %d %d %d \n", i, v.row, v.col, v.value)
		//fmt.Println(i, v.row, v.col, v.value)
		fw.WriteString(ss)
	}
	fw.Flush()
	fmt.Println("写入文件arr.txt成功！")
	//4、将文件读取出来直接打印原始数组
	rfile, err := os.Open("F:\\goland\\go_project\\21weeks\\Data_Structure\\sparsearray\\arr.txt")
	rf := bufio.NewReader(rfile)
	var arr = new([11][11]int)
	for {
		line, err := rf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}
		if err != nil {
			fmt.Printf("文件读取错误，%v\n", err)
			return
		}
		for i, v := range line {
			if i != 0 {
				// 把哪行哪列追加到数组中

			}
		}
	}

}

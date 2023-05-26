package main

import "fmt"

func setWay(myMap *[7][8]int, i, j int) bool {

	//什么情况下就可以找到出口
	if myMap[5][6] == 2 {
		return true
	} else { //开始找
		if myMap[i][j] == 0 {
			// 往下找试试
			myMap[i][j] = 2
			//每次都会开辟一个新的栈空间
			if setWay(myMap, i+1, j) {
				return true
			} else if setWay(myMap, i, j+1) {
				return true
			} else if setWay(myMap, i-1, j) {
				return true
			} else if setWay(myMap, i, j-1) {
				return true
			} else {
				myMap[i][j] = 3
				fmt.Println("没有通路")
				return false
			}
		} else {
			return false //此路为墙
		}
	}
}
func main() {
	// 创建迷宫
	// 0 没有走过的路
	// 1 一堵墙
	// 2 可以走
	// 3 走过但是没有走通

	var myMap [7][8]int
	for i := 0; i < 7; i++ {
		for j := 0; j < 8; j++ {
			myMap[i][0] = 1
			myMap[i][7] = 1
			myMap[0][j] = 1
			myMap[6][j] = 1
		}
	}

	myMap[3][1] = 1
	myMap[3][2] = 1
	for _, v := range myMap {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	setWay(&myMap, 1, 1)
	fmt.Println("探寻后的地图：")
	for _, v := range myMap {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}

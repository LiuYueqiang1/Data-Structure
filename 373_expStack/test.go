package main

import (
	"fmt"
	"strconv"
)

// 测试字符转换
func main() {
	exp := "3+3*4-3"
	ch := exp[0:1]
	ch1 := exp[3:4]
	//fmt.Println(ch)  3

	temp1, _ := strconv.Atoi(ch)
	temp3, _ := strconv.Atoi(ch1)
	//fmt.Println(temp1) //3
	fmt.Println(temp3) // * 但是确是 0
	temp2 := int([]byte(ch)[0])
	fmt.Println(temp2) //51   ASCiI码

	res := temp1 + temp3 + temp1

	fmt.Println("res:", res)
}

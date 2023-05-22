package main

import "fmt"

type numt struct {
	n    int
	next *numt
}

func add(x *numt) {
	tmp := x
	tmp = tmp.next
}
func main() {
	num := numt{}
	num1 := numt{}

	fmt.Println(num)
	fmt.Println(5 % 5)
}

package main

import "fmt"

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		// 插入的值为
		insertValue := arr[i]
		insertIndex := i - 1 //不是要插入的位置，而是要比较的位置
		for insertIndex >= 0 && arr[insertIndex] < insertValue {
			arr[insertIndex+1] = arr[insertIndex] //值往后移动一位
			//需要跟前面的值比较，不然比较的只是最后一位，直到循环条件不存在后结束，
			// 获取到此时index Index的值+1为插入的位置
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertValue
		}
		fmt.Printf("第 %d 次后插入后的数组为:%v\n", i, *arr)
	}
}
func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
}

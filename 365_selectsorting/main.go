package main

import "fmt"

// 选择排序
func selectsort(arr [6]int) {
	for j := 0; j < len(arr); j++ {
		//1、假设当前的就是最大值
		max := arr[j]
		maxIndex := j
		//2、将此值与剩下的值进行比较
		//找到真正的最大值
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 如果此时数组索引不等于最大值的索引，则交换
		if j != maxIndex {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次 %v\n", j+1, arr)
	}
}
func main() {
	arr := [6]int{5, 3, 55, 62, 1, 4}
	selectsort(arr)
}

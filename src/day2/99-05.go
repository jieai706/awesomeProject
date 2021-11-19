package main

import "fmt"

/**
输入三个整数x,y,z，请把这三个数由小到大输出
 */
func main() {
	fmt.Println(compareNumber(compareNumber(10, 4), 16))
}

func compareNumber(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
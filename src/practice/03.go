package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer语句1")
	fmt.Println("正常语句")
	defer fmt.Println("defer语句2")
	fmt.Println(func001())

	fmt.Println(calc(10, 20, add))
}

func func001() (a int, b float64) {
	a = 1
	b = 0.5
	return a,b
}

func add(a, b int) int {
	return a + b
}

func calc(a, b int,op func(int ,int) int) int {
	return op(a, b)
}

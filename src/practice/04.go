package main

import (
	"fmt"
	"strings"
)

func checkSuffix(suffix string) func(string) bool {
	return func(filename string) bool {
		if !strings.HasSuffix(filename, suffix) {
			return false
		}
		return true
	}
}

/**
判断是否以某个子串结尾
 */
func main() {
	check := checkSuffix(".txt")
	fmt.Println(check("test"))
	fmt.Println(check("hello.txt"))

	add, sub := test01(10)
	fmt.Println(add(1))
	fmt.Println(sub(3))
}

//返回 2 个闭包
func test01(base int) (func(int) int, func(int) int) {
	add := func(n int) int {
		base = base + n
		return base
	}

	sub := func(n int) int {
		base = base - n
		return base
	}
	return add, sub
}
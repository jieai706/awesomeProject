package main

import "fmt"

/**
struct
*/
func main() {
	p1 := Person{1, 18, "男", "张三", "13111111111"}
	fmt.Println(p1)
	fmt.Println(&p1)
	var p2 *Person = &Person{2, 20, "男", "李四", "13211111111"}
	fmt.Println(*p2)
	fmt.Println(&p2)
	fmt.Println(p2)
}

type Person struct {
	id   int
	age  int
	sex  string
	name string
	tel  string
}

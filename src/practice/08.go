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
	var stu1 = Student{
		Person: Person{3, 22, "男", "王五", "13311111111"},
		stuNo:  0,
		school: "whu",
	}
	fmt.Println(stu1)
}

type Person struct {
	id   int
	age  int
	sex  string
	name string
	tel  string
}

type Student struct {
	Person
	stuNo  int
	school string
}

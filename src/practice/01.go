package main

import (
	"fmt"
	"reflect"
)

const constV1 = 123

type myint int

func main () {
	var v1 int
	v1 = 1
	var v2 = 2
	v3 := 3
	v4, _ := test()
	fmt.Printf("%b\n", v3)
	fmt.Println(v1,v2,v3, constV1,v4)
	var v5 float64 = 1.4
	fmt.Println(reflect.TypeOf(v5))
	var v6 bool
	v6 = true
	v7 := !v6
	fmt.Println(v6,v7)
	v8 := `hello
asdfgf
	world
`
	fmt.Println(v8)
	var v9 myint
	fmt.Println(reflect.TypeOf(v9))
}

func test() (int, int) {
	return 4,5
}
package main

import "fmt"
import "reflect"

func main() {
	var a = 1.0
	fmt.Println(reflect.TypeOf(a))
	b := reflect.TypeOf(a).Kind()
	fmt.Println(b)
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(a).Kind() == reflect.Int)
}

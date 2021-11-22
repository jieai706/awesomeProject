package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)
	fmt.Printf("arr2: %v\n", arr)
	var sli []int = make([]int, 1, 2)
	var sli2 = []int{1, 3, 2, 4}
	var sli3 = []int{1, 3, 2, 4}
	fmt.Println(sli, sli2)
	fmt.Println(append(sli2, 2, 4))
	fmt.Println(append(sli2, sli3...))
}

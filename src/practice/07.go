package main

import "fmt"

/**
pointer
*/
func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
	swapValue(&a, &b)
	fmt.Println(a, b)
}

func swapValue(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b *int) (*int, *int) {
	a, b = b, a
	return a, b
}

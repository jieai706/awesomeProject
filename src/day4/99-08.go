package main

import "fmt"

/**
输出9*9口诀
*/
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i >= j {
				fmt.Printf("%d*%d=%d ", i, j, i*j)
			}
		}
		fmt.Println()
	}
}

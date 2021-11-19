package main

import (
	"fmt"
)

func main() {
	str := "a"
	switch str {
	case "a":
		fmt.Println("a")
		fallthrough
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("c")
	}

	//for {
	//	fmt.Println("wait sleep")
	//	time.Sleep(0 * time.Second)
	//}

	myArray := [4]int{1,2,3,4}
	for i := range myArray {
		fmt.Printf("index: %d, value: %d\n", i, myArray[i])
	}

	for i,v := range myArray {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

}



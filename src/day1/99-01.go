package main

import "fmt"

/**
题目：有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
 */
func main() {
	myArray := [4]int{1,2,3,4}
	fmt.Println("生成的总数为：", day01func(myArray))
}

func day01func(myArray [4]int) (result int) {
	//for i := 0;i < len(myArray);i ++ {
	//	for j :=0;j < len(myArray);j ++ {
	//		for k :=0;k < len(myArray);k ++ {
	//			fmt.Println("生成的三位数是", myArray[i], myArray[j], myArray[k])
	//			result ++
	//		}
	//	}
	//}
	//for i := range myArray {
	//	for j := range myArray {
	//		for k := range myArray {
	//			fmt.Println("生成的三位数是", myArray[i], myArray[j], myArray[k])
	//			result ++
	//		}
	//	}
	//}
	for _,vi := range myArray {
		for _,vj := range myArray {
			for _,vk := range myArray {
				fmt.Println("生成的三位数是", vi, vj, vk)
				result ++
			}
		}
	}
	return result
}

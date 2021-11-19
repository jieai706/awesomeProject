package main

import "fmt"

/**
输入某年某月某日，判断这一天是这一年的第几天
 */
func main() {
	fmt.Println(calcDay(2021,11,18))
}

func calcDay(year int, month int, day int) int {
	normalYearArray := [12]int{31,29,31,30,31,30,31,31,30,31,30,31}
	specialYearArray := [12]int{31,28,31,30,31,30,31,31,30,31,30,31}

	var result int = day

	// 判断是否是闰年
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		for i := 0;i < month - 1;i ++ {
			result += normalYearArray[i]
		}
	} else {
		for i := 0;i < month - 1;i ++ {
			result += specialYearArray[i]
		}
	}

	return result
}

func calcDay2(year int, month int, day int) int {
	var sum, feb int
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		feb = 29
	} else {
		feb = 28
	}
	switch month {
	case 1:

		sum = 0 + day

	case 2:

		sum = 31 + day

	case 3:

		sum = feb + 31 + day

	case 4:

		sum = 31 + feb + 31 + day

	case 5:

		sum = 30 + 31 + feb + 31 + day

	case 6:

		sum = 31 + 30 + 31 + feb + 31 + day

	case 7:

		sum = 30 + 31 + 30 + 31 + feb + 31 + day

	case 8:

		sum = 31 + 30 + 31 + 30 + 31 + feb + 31 + day

	case 9:

		sum = 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day

	case 10:

		sum = 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day

	case 11:

		sum = 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day

	case 12:

		sum = 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31 + feb + 31 + day

	}

	return sum
}

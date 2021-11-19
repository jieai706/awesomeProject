package main

import "fmt"
/**
利润(I)低于或等于10万元时，奖金可提10%；
利润高于10万元，低于20万元时，低于10万元的部分按10%提成，高于10万元的部分，可提成7.5%；
20万到40万之间时，高于20万元的部分，可提成5%；40万到60万之间时高于40万元的部分，可提成3%；
60万到100万之间时，高于60万元的部分，可提成1.5%；
高于100万元时，超过100万元的部分按1%提成；
从键盘输入当月利润，求应发放奖金总数？
 */
func main() {
	var bonus float64
	//a := 20000
	fmt.Scan(&bonus)
	fmt.Println(bonus)
	fmt.Println(day01func2(bonus))
}

func day01func2(bonus float64) float64 {
	var result float64
	if bonus <= 100000 {
		result = bonus * 0.1
	} else if bonus <= 200000 {
		result = (bonus - 100000) * 0.075 + 100000 * 0.1
	} else if bonus <= 400000 {
		result = (bonus - 200000) * 0.05 + 100000 * 0.1 + 200000 * 0.075
	} else if bonus <= 600000 {
		result = (bonus - 400000) * 0.03 + 100000 * 0.1 + 200000 * 0.075 + 200000 * 0.05
	} else if bonus <= 1000000 {
		result = (bonus - 600000) * 0.015 + 100000 * 0.1 + 200000 * 0.075 + 200000 * 0.05 + 200000 * 0.03
	} else {
		result = (bonus - 1000000) * 0.01 + 100000 * 0.1 + 200000 * 0.075 + 200000 * 0.05 + 200000 * 0.03 + 400000 * 0.015
	}
	return result
}

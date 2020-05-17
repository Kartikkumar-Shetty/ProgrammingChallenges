package main

import "fmt"

func main() {
	count := 0
	for i := 9; i > 0; i-- {
		for j := 9; j >= 0; j-- {
			for k := 9; k >= 0; k-- {
				num := (i * 100000) + (j * 10000) + (k * 1000) + (k * 100) + (j * 10) + (i)
				if num > 998001 {
					continue
				}
				isFactor, _, _ := checkNumberHas3DigitFactor(num)
				if isFactor {
					fmt.Println(num)
					return
				}
				count = count + 1
			}
		}
	}
}
func checkNumberHas3DigitFactor(number int) (bool, int, int) {
	var qt int
	for i := 1000; i >= 100; i-- {
		if number%i == 0 {
			qt = number / i
			if qt > 99 && qt < 1000 {
				return true, i, qt
			}
		}
	}
	return false, 0, 0
}

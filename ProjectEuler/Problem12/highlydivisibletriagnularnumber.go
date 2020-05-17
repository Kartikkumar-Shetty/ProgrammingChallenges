package main

import "fmt"

func main() {
	trianglenumber()
}

func trianglenumber() {
	sum := 0
	i := 1
	for {
		sum = sum + i
		numOfDiv := numOfDivisors(sum)
		if numOfDiv > 500 {
			fmt.Println(sum)
			break
		}
		// //fmt.Println(sum, ":", numOfDiv)
		// if i > 10 {
		// 	break
		// }
		i++
	}
}

func numOfDivisors(num int) int {
	max := num
	numofFactors := 0
	for i := 1; i < max; i++ {
		if num%i == 0 {
			max = num / i
			numofFactors = numofFactors + 2
		}

	}
	return numofFactors
}

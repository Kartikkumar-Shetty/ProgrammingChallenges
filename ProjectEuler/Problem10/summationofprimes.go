package main

import "fmt"

func main() {
	primes()
}

func primes() {
	n := 2000000
	var arr [2000001]int
	var start int

	for i := 2; i <= n; i++ {

		if arr[i] == 1 {
			continue
		}
		start = i

		squarevalue := start * start
		if squarevalue > n {
			break
		}

		for j := squarevalue; j <= n; j++ {
			if j%start == 0 {
				arr[j] = 1
			}
		}

	}

	sum := 0
	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}

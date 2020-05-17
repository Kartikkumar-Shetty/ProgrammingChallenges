package main

import "fmt"

func main() {
	largestfactor := primefactors(600851475143, 0)
	fmt.Println(largestfactor)
}
func primefactors(num, largestfactor int) int {
	for i := 2; i < num; i++ {
		if checkifPrime(i) {
			if num%i == 0 {
				if largestfactor < i {
					largestfactor = i
					fmt.Println(largestfactor, " ", num/i, " ", num%i)
					return primefactors(num/i, largestfactor)
				} else {
					continue

				}
			}
		}
	}
	return largestfactor
}

func checkifPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

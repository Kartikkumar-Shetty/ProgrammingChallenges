package main

import "fmt"

func main() {

	for i := 1; i < 10000; i++ {

		arr1 := divisors(i)
		sum1 := sumOfDivisors(arr1)
		if sum1 > 10000 {
			continue
		}

		arr2 := divisors(sum1)
		sum2 := sumOfDivisors(arr2)

		if i == sum2 {
			fmt.Println(i, " ", sum1, ":", sum2)
		}

	}
}
func divisors(num int) []int {
	arr := []int{1}
	max := num - 1
	for i := 2; i < max; i++ {
		if num%i == 0 {
			max = num / i
			arr = append(arr, i)
			arr = append(arr, num/i)
		}
	}
	return arr

}

func sumOfDivisors(div []int) int {
	sum := 0
	for i := 0; i < len(div); i++ {
		sum = sum + div[i]
	}
	return sum
}

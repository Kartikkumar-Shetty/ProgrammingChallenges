package main

import (
	"fmt"
)

var counter uint64
var sum uint64
var k uint64

func main() {
	fibonacci()
}

func fibonacci() {
	var prev, curr, next uint64
	curr = 2
	prev = 1
	sum = 2
	for {
		next = prev + curr
		prev = curr
		curr = next
		if next >= 4000000 {
			break
		}
		if next%2 == 0 {
			sum = sum + next
		}

	}
	fmt.Println(sum)
}

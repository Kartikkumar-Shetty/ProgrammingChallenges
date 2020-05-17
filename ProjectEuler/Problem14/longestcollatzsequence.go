package main

import "fmt"

//we will try reversing the sequence
//n->n/2 where n is even
//n->3n+1 where n is odd

//if n is even then 2n
//if n is of the form 3n+1 then-> if n is even then either 2n or (n-1)/3 && if n is odd then (n-1)/3
//

type node struct {
	val  int
	next *node
}

var exists bool
var branch []*node
var first *node

func main() {
	//branchcounter := 0
	var greatest int
	for i := 1000000; i > 2; i-- {
		len := createSequence(i)
		if len > greatest {
			greatest = len
			fmt.Println(greatest, " - ", i)

		}
	}
}

func createSequence(n int) int {
	//fmt.Println("")
	sequencelength := 1
	for {
		//	fmt.Print(n, " ")
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		sequencelength++

	}
	return sequencelength
}

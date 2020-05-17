package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100.0

	sumofsquare := (n * (n + 1) * (2*n + 1)) / 6
	squareofsum := math.Pow(((n * (n + 1)) / 2), 2)
	fmt.Println(int(squareofsum - sumofsquare))
}

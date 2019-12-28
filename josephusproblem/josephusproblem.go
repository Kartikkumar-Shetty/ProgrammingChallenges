package main

import "fmt"

func main() {
	line := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	outerloopcount := 0
	for i := 0; i <= len(line); i++ {
		if i >= len(line) {
			i = 0
		}
		if line[i] == 0 {
			outerloopcount++
			if outerloopcount > len(line) {
				break
			}
			continue
			outerloopcount = 0
		}
		fmt.Println("i:", i)
		loopcount := 0
		done := false
		for j := i + 1; j <= len(line); j++ {
			loopcount++
			if loopcount >= len(line) {
				done = true
				fmt.Println("done")
				break

			}
			if j >= len(line) {
				j = 0
			}
			fmt.Println("j:", j)
			if line[j] != 0 {
				fmt.Printf("%v \n", line)
				line[j] = 0
				i++
				fmt.Println("killed: loopcount: ", loopcount)
				break
			}
		}
		if i >= len(line) {
			i = 0
		}
		if done {
			fmt.Println("Done")
			break
		}
		outerloopcount = 0

	}
	for i, r := range line {
		if r != 0 {
			fmt.Print("Position:", i+1, "; Value: ", r)
			fmt.Printf("%v \n", line)
		}
	}
}

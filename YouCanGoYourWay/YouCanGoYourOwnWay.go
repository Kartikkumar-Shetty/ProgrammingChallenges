package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var mdim int
	var totalCases int
	var givenDirections string

	scanf("%d\n", &totalCases)
	caseNum := 1
	for {
		if caseNum > totalCases {
			return
		}
		scanf("%d\n", &mdim)

		// maze := createMatrix(mdim)
		// fmt.Println(maze)
		scanf("%s\n", &givenDirections)
		arrDirections := strings.Split(givenDirections, "")
		//	fmt.Println("herDir", arrDirections)

		givenSol := createPath(arrDirections, mdim)
		mySol := Walk(givenSol, arrDirections, mdim)
		output := fmt.Sprintf("%s %s%d: %s\n", "Case", "#", caseNum, mySol)
		fmt.Print(output)
		caseNum++
		continue

	}
}

func createPath(arrDirections []string, dim int) map[int]int {
	sol := make(map[int]int)
	row := 0
	col := 0
	for _, val := range arrDirections {
		if val == "S" {
			sol[(row*dim)+col] = ((row + 1) * dim) + col
			row++

		} else {
			sol[(row*dim)+col] = (row * dim) + (col + 1)
			col++

		}

	}
	return sol

}

func createMatrix(dim int) []int {
	matrix := make([]int, dim*dim)
	for i, _ := range matrix {
		matrix[i] = i
	}
	return matrix
}

func Walk(givenSol map[int]int, arrDirections []string, dim int) string {
	myDirections := ""
	row := 0
	col := 0
	//	fmt.Println(givenSol)
	for _, v := range arrDirections {
		// fmt.Println("She was going South:", v)
		// fmt.Println(v)
		switch strings.Trim(strings.ToUpper(v), " ") {
		case "S":
			// fmt.Println("She was going South")
			// fmt.Println("(row*dim)+col, ((row+1)*dim)+col", (row*dim)+col, ((row+1)*dim)+col)

			if givenSol[(row*dim)+col] == ((row+1)*dim)+col {
				myDirections = fmt.Sprintf("%s%s", myDirections, "E")
				//				fmt.Println(("going east"))
				col++
				continue
			}

			//fmt.Println("(row*dim)+col == (row*dim)+(col+1)", (row*dim)+col, (row*dim)+(col+1))
			if givenSol[(row*dim)+col] == (row*dim)+(col+1) {
				myDirections = fmt.Sprintf("%s%s", myDirections, "S")
				//				fmt.Println(("going south"))
				row++
				continue
			}

			myDirections = fmt.Sprintf("%s%s", myDirections, "E")
			//			fmt.Println(("going east"))
			col++
			continue

		case "E":
			// fmt.Println("She was going East")
			// fmt.Println("(row*dim)+col == (row*dim)+(col+1)", (row*dim)+col, (row*dim)+(col+1))
			if givenSol[(row*dim)+col] == (row*dim)+(col+1) {
				myDirections = fmt.Sprintf("%s%s", myDirections, "S")
				// fmt.Println(("going south"))
				row++
				continue
			}
			//fmt.Println("(row*dim)+col == ((row+1)*dim)+col", (row*dim)+col, ((row+1)*dim)+col)
			if givenSol[(row*dim)+col] == ((row+1)*dim)+col {
				myDirections = fmt.Sprintf("%s%s", myDirections, "E")
				// fmt.Println(("going East"))
				col++
				continue
			}
			myDirections = fmt.Sprintf("%s%s", myDirections, "S")
			// fmt.Println(("going south"))
			row++
			continue
		}

	}
	return myDirections
}

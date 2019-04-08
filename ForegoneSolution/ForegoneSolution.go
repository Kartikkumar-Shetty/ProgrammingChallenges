package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

//func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var input string
	var totalCases int

	scanf("%d\n", &totalCases)
	caseNum := 1
	for {
		if caseNum > totalCases {
			return
		}
		scanf("%s\n", &input)
		//fmt.Println(input)
		input = strings.Replace(input, "\n", "", -1)
		if input == "" {
			return
		}

		if strings.Contains(input, "4") {
			arrNums := strings.Split(input, "")
			fourPos := FindPositionOfFour(arrNums)
			for _, val := range fourPos {
				num := arrNums[val]
				intNum, _ := strconv.Atoi(num)
				arrNums[val] = fmt.Sprintf("%d+1", intNum-1)
			}
			//fmt.Println(arrNums)
			num1 := 0
			num2 := 0
			for i, _ := range arrNums {
				if strings.Contains(arrNums[i], "+") {
					arrVals := strings.Split(arrNums[i], "+")
					intVal1, _ := strconv.Atoi(arrVals[0])
					intVal2, _ := strconv.Atoi(arrVals[1])
					//fmt.Println(num1, intVal1, "*", "expo", 10, (len(arrNums)-1)-i)
					//fmt.Println(num1)
					num1 = num1 + (intVal1 * expo(10, ((len(arrNums)-1)-i)))
					//fmt.Println(num2, intVal2, "*", "expo", 10, (len(arrNums)-1)-i)
					num2 = num2 + (intVal2 * expo(10, ((len(arrNums)-1)-i)))
					//fmt.Println(num2)
					continue
				}
				intVal, _ := strconv.Atoi(arrNums[i])
				//fmt.Println(num1, intVal, "*", "expo", 10, (len(arrNums)-1)-i)
				num1 = num1 + (intVal * expo(10, ((len(arrNums)-1)-i)))
				//fmt.Println(num1)
			}
			output := fmt.Sprintf("%s %s%d: %d %d\n", "Case", "#", caseNum, num1, num2)
			fmt.Print(output)
			caseNum++
			continue
		}
		print("")
		caseNum++
	}
	writer.Flush()

}

func FindPositionOfFour(input []string) []int {
	var pos []int
	for i, val := range input {
		if val == "4" {
			pos = append(pos, i)
		}
	}
	return pos
}

func expo(x, y int) int {

	i := float64(x)
	j := float64(y)
	r := math.Pow(i, j)
	return int(r)
}

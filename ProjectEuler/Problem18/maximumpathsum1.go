package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum int
var globalcounter int

type node struct {
	val   int
	rside *node
	left  *node
	right *node
}

type stack struct {
	arr []int
}

func (s *stack) pop() {
	s.arr = s.arr[0:(len(s.arr) - 1)]
	//s.arr = s.arr[0:2]
}
func (s *stack) push(n int) {
	s.arr = append(s.arr, n)
}
func (s *stack) print() {
	for _, v := range s.arr {
		fmt.Print(v, " ")
	}
}

func main() {
	root := initTree()
	//fmt.Println("-------------------------------------------")
	//parseTree(root)
	//printTree(root)
	createStack(root, stack{})
	fmt.Println("Done")
	fmt.Println(sum)
	fmt.Println(globalcounter)
}
func createStack(n *node, s stack) {
	s.push(n.val)
	if n.left != nil {
		createStack(n.left, s)
	}
	if n.right != nil {
		createStack(n.right, s)
	}
	if n.left == nil && n.right == nil {
		checkSum(s.arr)
	}

	s.pop()
	//fmt.Print(n.val, " ")
}
func checkSum(arr []int) {
	globalcounter++
	for i := 0; i < len(arr)-1; i++ {
		temp := sumOfarr(arr)
		if temp > sum {
			fmt.Println(arr)
			sum = temp
			fmt.Println(sum)
		}
	}
}

func sumOfarr(arr []int) int {
	s := 0
	for _, v := range arr {
		s = s + v
	}
	return s
}

func parseTree(n *node) {
	if n.left != nil {
		fmt.Print(n.val, "-l->")
		parseTree(n.left)
	}
	if n.right != nil {
		fmt.Print(n.val, "-r->")
		parseTree(n.right)
	}

	fmt.Print(n.val, " ")
	return
}

func printTree(left *node) {
	for {
		n := left

		for {
			fmt.Print(n.val, " ")
			if n.rside == nil {
				break
			}
			n = n.rside
		}
		if n.left == nil {
			break
		}
		left = left.left

		fmt.Println("")
	}

}

func initTree() *node {
	//var val int
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var root *node
	reader := bufio.NewReader(file)
	ipRow := 0
	var prevrow []*node
	for {

		text, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		rowVals := strings.Split(text, " ")

		prevrow = func(rowVals []string, prevrow []*node) []*node {
			var prev *node
			var currRow []*node
			j := 0

			for i := 0; i < len(rowVals); i++ {
				v, err := strconv.Atoi(strings.Replace(rowVals[i], "\n", "", -1))
				if err != nil {
					panic(err)
				}

				curr := &node{
					val: v,
				}
				if ipRow == 0 && i == 0 {
					root = curr
				}

				if len(prevrow) > 0 {
					if i != 0 {
						prevrow[j].right = curr
						if j < len(prevrow)-1 {
							prevrow[j+1].left = curr
						}

						j++
					} else {
						prevrow[j].left = curr
					}

				}
				if i == 0 {
					prev = curr
				} else {
					prev.rside = curr
					prev = curr
				}
				currRow = append(currRow, curr)

			}
			return currRow
		}(rowVals, prevrow)

		// for _, r := range prevrow {
		// 	fmt.Print(r.val, " ")
		// }
		// fmt.Println("")

		ipRow++
	}
	return root
}

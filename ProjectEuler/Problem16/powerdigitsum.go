package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var prod uint64

func main() {
	i := big.NewInt(2)
	fmt.Println(i)
	num := fmt.Sprint(new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil))
	sumOfDigits(num)
}

func sumOfDigits(num string) {
	arr := strings.Split(num, "")
	sum := 0
	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		sum = sum + num
	}
	fmt.Print(" ", sum)
}

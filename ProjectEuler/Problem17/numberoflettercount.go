package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 999; i++ {
		if i < 10 {
			fmt.Print(getUnitsPlace(i))
		} else if i > 9 && i < 100 {
			fmt.Print(handleTwoDigit(i))
		} else if i > 99 && i < 1000 {
			splitval := strings.Split(fmt.Sprint(i), "")

			units, _ := strconv.Atoi(splitval[2])
			tens, _ := strconv.Atoi(splitval[1])
			hundreds, _ := strconv.Atoi(splitval[0])

			num := getHundredsPlace(hundreds)

			if tens == 0 && units != 0 {
				num = fmt.Sprint(num, "and", getUnitsPlace(units))
			} else if tens == 1 {
				num = fmt.Sprint(num, "and", get10to19(10+units))
			} else if tens > 1 {
				num = fmt.Sprint(num, "and", getTensPlace(tens), getUnitsPlace(units))
			}
			fmt.Print(num)

		}
	}
	fmt.Print("onethousand")
}

func handleTwoDigit(i int) string {
	if i > 9 && i < 20 {
		return fmt.Sprint(get10to19(i))
	} else if i > 19 && i < 100 {
		splitval := strings.Split(fmt.Sprint(i), "")
		units, _ := strconv.Atoi(splitval[1])
		tens, _ := strconv.Atoi(splitval[0])
		num := getTensPlace(tens)
		return fmt.Sprint(num, getUnitsPlace(units))
	}
	return ""
}

func getUnitsPlace(n int) string {
	num := ""
	switch n {
	case 1:
		num = "one"
	case 2:
		num = "two"
	case 3:
		num = "three"
	case 4:
		num = "four"
	case 5:
		num = "five"
	case 6:
		num = "six"
	case 7:
		num = "seven"
	case 8:
		num = "eight"
	case 9:
		num = "nine"
	}
	return num
}

func getTensPlace(n int) string {
	num := ""
	switch n {
	case 2:
		num = "twenty"
	case 3:
		num = "thirty"
	case 4:
		num = "forty"
	case 5:
		num = "fifty"
	case 6:
		num = "sixty"
	case 7:
		num = "seventy"
	case 8:
		num = "eighty"
	case 9:
		num = "ninety"
	}
	return num
}

func get10to19(n int) string {
	num := ""
	switch n {
	case 10:
		num = "ten"
	case 11:
		num = "eleven"
	case 12:
		num = "twelve"
	case 13:
		num = "thirteen"
	case 14:
		num = "fourteen"
	case 15:
		num = "fifteen"
	case 16:
		num = "sixteen"
	case 17:
		num = "seventeen"
	case 18:
		num = "eighteen"
	case 19:
		num = "nineteen"
	}
	return num
}
func getHundredsPlace(n int) string {
	num := ""
	switch n {
	case 1:
		num = "onehundred"
	case 2:
		num = "twohundred"
	case 3:
		num = "threehundred"
	case 4:
		num = "fourhundred"
	case 5:
		num = "fivehundred"
	case 6:
		num = "sixhundred"
	case 7:
		num = "sevenhundred"
	case 8:
		num = "eighthundred"
	case 9:
		num = "ninehundred"
	}
	return num
}

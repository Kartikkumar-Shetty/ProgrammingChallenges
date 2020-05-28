package main

import (
	"fmt"
	"time"
)

func main() {
	layoutISO := "2006-01-02"
	day, err := time.Parse(layoutISO, "1901-01-01")
	if err != nil {
		panic(nil)
	}
	numOfSunday := 0
	for {
		if day.Year() == 2001 {
			break
		}

		if day.Weekday() == time.Sunday {
			numOfSunday++
		}
		day = day.AddDate(0, 1, 0)

	}
	fmt.Println(numOfSunday)
}

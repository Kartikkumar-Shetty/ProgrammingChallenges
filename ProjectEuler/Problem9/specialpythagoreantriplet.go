package main

import (
	"fmt"
	"math"
)

func main() {
	//a=m^2+n^2
	//b=2mn
	//c=m^2-n^2

	//m(m+n)=500

	//hence m, m+n
	//5,100  25,20  125,4  50,10

	//possible values  m=5 m+n=5+95, m=20 m+n=20+5, m=4 m+n = 4+121, m=10 m+n=10+40

	fmt.Println(math.Pow(5, 2)+math.Pow(95, 2), " ", 2*5*95, " ", math.Pow(5, 2)-math.Pow(95, 2))
	fmt.Println(math.Pow(20, 2)+math.Pow(5, 2), " ", 2*20*5, " ", math.Pow(20, 2)-math.Pow(5, 2))
	fmt.Println(math.Pow(4, 2)+math.Pow(121, 2), " ", 2*4*121, " ", math.Pow(4, 2)-math.Pow(121, 2))
	fmt.Println(math.Pow(10, 2)+math.Pow(40, 2), " ", 2*10*40, " ", math.Pow(10, 2)-math.Pow(40, 2))

}

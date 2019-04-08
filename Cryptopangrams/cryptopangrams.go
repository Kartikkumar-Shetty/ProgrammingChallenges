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

	var totalCases int
	scanf("%d\n", &totalCases)
	caseNum := 1

	for {
		allAplhabets := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
		var primeMax int
		cipherFactors := make(map[int]int)
		var valCount int
		var cipher []int
		var temp int
		finalMap := make(map[int]string)
		var output string
		var messageOrder []int
		if caseNum > totalCases {
			return
		}

		scanf("%d %d\n", &primeMax, &valCount)

		primes := getAllPrimes(primeMax)
		for i := 0; i <= valCount-1; i++ {
			if i == valCount-1 {
				scanf("%d\n", &temp)
			} else {
				scanf("%d", &temp)
			}
			cipher = append(cipher, temp)
		}
		oldp1 := 0
		oldp2 := 0
		p2 := 0
		for i, c := range cipher {
			for _, p1 := range primes {
				if c%p1 == 0 {
					cipherFactors[p1] = p1
					p2 = c / p1
					cipherFactors[p2] = p2

					if oldp1 == p1 {
						messageOrder = append(messageOrder, oldp2)
						if i == len(cipher)-1 {
							messageOrder = append(messageOrder, oldp1)
							messageOrder = append(messageOrder, p2)
						}
					} else if oldp1 == p2 {
						messageOrder = append(messageOrder, oldp2)
						if i == len(cipher)-1 {
							messageOrder = append(messageOrder, oldp1)
							messageOrder = append(messageOrder, p1)
						}
					} else if oldp2 == p1 {
						messageOrder = append(messageOrder, oldp1)
						if i == len(cipher)-1 {
							messageOrder = append(messageOrder, oldp2)
							messageOrder = append(messageOrder, p2)
						}
					} else if oldp2 == p2 {
						messageOrder = append(messageOrder, oldp1)
						if i == len(cipher)-1 {
							messageOrder = append(messageOrder, oldp2)
							messageOrder = append(messageOrder, p1)
						}
					}

					oldp1 = p1
					oldp2 = p2

					break
				}
			}
		}

		factors := Sort(getKeys(cipherFactors))

		for i, f := range factors {
			finalMap[f] = allAplhabets[i]

		}
		for _, c := range messageOrder {
			output = fmt.Sprintf("%s%s", output, finalMap[c])
		}

		output = fmt.Sprintf("%s%d: %s\n", "Case #", caseNum, output)
		fmt.Print(output)
		caseNum++
	}

}

func getKeys(mp map[int]int) []int {
	val := []int{}
	for k, _ := range mp {
		val = append(val, k)
	}
	return val

}

func getAllPrimes(n int) []int {
	prime := []int{}
	sieve := make(map[int]int)
	i := 2
	for i <= n {
		sieve[i] = 0
		i++
	}

	i = 2
	var num int
	for i <= n {
		num = i
		j := num * num
		for j <= n {
			if j%num == 0 {
				sieve[j] = 1
			}
			j++
		}
		i++

	}
	i = 2
	for i <= n {
		if sieve[i] == 0 {
			prime = append(prime, i)
		}
		i++
	}
	return prime

}

func Sort(factors []int) []int {

	if len(factors) < 2 {
		return factors
	}
	m := (len(factors)) / 2
	return Merge(Sort(factors[:m]), Sort(factors[m:]))
}

func Merge(l, r []int) []int {

	size, i, j := len(l)+len(r), 0, 0
	fact := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(l)-1 && j <= len(r)-1 {
			fact[k] = r[j]
			j++
		} else if j > len(r)-1 && i <= len(l)-1 {
			fact[k] = l[i]
			i++
		} else if l[i] < r[j] {
			fact[k] = l[i]
			i++
		} else {
			fact[k] = r[j]
			j++
		}
	}
	return fact
}

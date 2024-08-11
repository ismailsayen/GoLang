package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		return
	}

	nb, err := strconv.Atoi(args[0])
	if nb == 0 {
		return
	}
	if nb == 1 {
		fmt.Println()
		return
	}

	if err != nil {
		return
	}
	if isPrime(nb) {
		fmt.Println(nb)
		return
	}
	fmt.Println(findPrimeFactors(nb))
}

func findPrimeFactors(nb int) string {
	res := ""
	res2 := ""
	for i := 2; i <= nb; i++ {
		if  isPrime(i) && nb%i == 0 {
			nb /= i
			res += Atoi(i) + "*"
			i--
		}
	}
	for i := 0; i < len(res)-1; i++ {
		res2 += string(res[i])
	}
	return res2
}

func isPrime(a int) bool {
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(a int) string {
	res := ""
	for a > 0 {
		res = string('0'+a%10) + res
		a /= 10
	}
	return res
}

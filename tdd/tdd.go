package main

import (
	"fmt"
)

func isPrime(number int) bool {
	// sqrtRoot := math.Sqrt(float64(number))

	switch number {
	case 0, 1:
		return false
	}

	for j := 2; j < number; j++ {
		if number%j == 0 {
			return false
		}
	}
	return true
}

func primeGen(count int) []int {
	primes := []int{}
	number := 0
	for len(primes) < count {
		if isPrime(number) {
			primes = append(primes, number)
		}
		number++
	}

	return primes
}

func main() {
	fmt.Println(primeGen(10))
}

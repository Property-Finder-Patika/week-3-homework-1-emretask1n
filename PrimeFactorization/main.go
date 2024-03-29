package main

import "fmt"

// Getting all prime factors of a given number n
func PrimeFactors(n int) (result []int) {

	// Get the number of 2s that divide n
	for n%2 == 0 {
		result = append(result, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			result = append(result, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		result = append(result, n)
	}

	return result
}

func main() {
	fmt.Println(PrimeFactors(12))
	fmt.Println(PrimeFactors(24))
	fmt.Println(PrimeFactors(57))
}

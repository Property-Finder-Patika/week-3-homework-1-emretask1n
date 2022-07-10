package main

import (
	"fmt"
)

func display(num int, prime []bool) {

	fmt.Printf("Primes less than %d : ", num)
	for i := 2; i <= num; i++ {
		if prime[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func findPrimes(num int) {

	// an array of boolean
	prime := make([]bool, num+1)

	// initialize everything with false first(not crossed)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true // cross
			}
		}

	}

	display(num, prime)

}

func main() {
	findPrimes(57)
	findPrimes(48)
}

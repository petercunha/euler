// Largest prime factor
// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143?

package main

import (
	"fmt"
	"math"
)

const NUMBER int = 600851475143

func main() {

	// The largest prime factor of a number is always less 
	// than the square root of said number, so we set the
	// start point of our loop at the square root of 600851475143
	n := int(math.Sqrt(float64(NUMBER)))

	// To expediate the process, we decrement from n rather than
	// increment from 2
	for n > 1 {
		if NUMBER % n == 0 {
			if isPrime(n) {
				fmt.Println("Highest prime:", n)
				break
			}
		}
		n--
	}
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}
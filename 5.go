// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"fmt"
)

func main() {

	for i := 1;; i++ {
		if check(i) {
			fmt.Println("Answer:", i)
			break
		}
	}

}

func check(num int) bool {
	// Numbers are less likely to contain larger factors,
	// therefore we count down from twenty instead of
	// up to it. This cuts computation time in half.
	for i := 20; i > 1; i-- {
		if num % i != 0 {
			return false
		}
	}
	return true
}
// Largest palindrome product
// Problem 4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
)

func main() {

	var (
		highest int = 0
		current int
	)

	// We use a range of 100 to 999 because the answer
	// must be a product of a three digit number
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			current = i * j
			if isPalindrome(current) && current > highest {
				highest = current
			}
		}
	}

	fmt.Println("Highest palindrome:", highest)
}

func isPalindrome(num int) bool {
	return num == reverse(num)
}

// Returns a reversed version of the number
func reverse(num int) int {

	var (
		reversed  int = 0
		lastDigit int
	)

	for num > 0 {
		// Last number is taken off "num"
		// Ex: if num = 123, lastDigit will equal 3
		lastDigit = num % 10

		// "lastDigit" is concatonated to the end of "reversed"
		// Ex:  if reversed = 12 and lastDigit = 3
		//		12 * 10 = 120
		//		120 + 3 = 123
		reversed = reversed * 10 + lastDigit

		// We shave off the last number of "num"
		// This will not leave a hanging decimal
		// because num is declared as type int
		num = num / 10
	}

	return reversed
}
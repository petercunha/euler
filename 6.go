package main

import "fmt"

const LIM = 100

func main() {
	s1, s2 := 0, 0
	for i := 0; i <= LIM; i++ {
		s1 += i*i; s2 += i
	}
	s2 *= s2
	fmt.Println("Result:", s2, "-", s1, "=", s2 - s1)
}
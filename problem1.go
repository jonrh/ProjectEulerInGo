package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(problem1())
}

// Problem description: If we list all the natural numbers below 10 that are
// multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is
// 23. Find the sum of all the multiples of 3 or 5 below 1000.
func problem1() string {
	sum := 0
	for i := 1; i < 1000; i++ {
		if isMultipleOf3or5(i) {
			sum = sum + i
		}
	}
	return strconv.Itoa(sum)
}

// Takes in an integer and returns true if it can be divided by 3 or 5.
func isMultipleOf3or5(number int) bool {
	if number%3 == 0 || number%5 == 0 {
		return true
	}
	return false
}

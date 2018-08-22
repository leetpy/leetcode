package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digit := []int{}
	for {
		if x < 10 {
			digit = append(digit, x)
			break
		}
		d := x % 10
		x /= 10
		digit = append(digit, d)
	}

	for i, v := range digit {
		if v != digit[len(digit)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}

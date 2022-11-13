package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// convert int to string
	xS := strconv.Itoa(x)
	// number of xS index
	nS := len(xS) - 1

	// iterates the string, to check it's palindrome or not
	// e.g 121 -> 1 match 1 then palindrome
	// 10 -> 1 not match 0, then it's not palindrome
	// -121 -> - not match with 1, then it's not palindrome
	for i, v := range xS {
		// if current index char from left
		// not match with the char from right (same index from backward)
		// return false as it's not palindrome
		if string(v) != string(xS[nS-i]) {
			return false
		}
	}

	// return true as it's palindrome
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}

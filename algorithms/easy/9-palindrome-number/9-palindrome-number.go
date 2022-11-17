package main

import (
	"fmt"
	"strconv"
)

func isPalindrome2(x int) bool {
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

func isPalindrome(x int) bool {
	// under 0 -> -1, -121 should not palindrome
	if x < 0 {
		return false
	}

	// single digit (1-9) must be palindrome
	if x < 10 {
		return true
	}

	// copy the input value to original variable
	original := x
	// prepare the reserve variable
	reserve := 0

	// do loop until x is 0
	for x != 0 {
		// reserve the input digit by digit from backward
		// for the example input 123
		// first  -> reserve*10 : 0,   x%10 : 3, reserve : 3,   x divide 10 : 12
		// second -> reserve*10 : 30,  x%10 : 2, reserve : 32,  x divide 10 : 1
		// third  -> reserve*10 : 320, x%10 : 1, reserve : 321, x divide 10 : 0
		// break the loop since x = 0
		reserve = reserve*10 + x%10

		// divide x by 10
		x /= 10
	}

	// check if original and reserve is same number
	return original == reserve
}

func main() {
	fmt.Println(isPalindrome(123))
}

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome2(x int) bool {
	// convert int to string
	xS := strconv.Itoa(x)
	// define the min and max index of xS
	// which always from 0 and the max is len-1
	i, j := 0, len(xS)-1

	// loop and stop if i more than j
	for i < j {
		// if current index char from left
		// not match with the char from right (same index from backward)
		// return false as it's not palindrome
		if xS[i] != xS[j] {
			return false
		}

		i++
		j--
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
	fmt.Println(isPalindrome(121))
}

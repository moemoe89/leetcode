package main

import (
	"fmt"
)

func isValid(s string) bool {
	// map for bracket's pair
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// stack for storing the open bracket
	stack := make([]rune, 0)

	for _, bracket := range s {
		// if bracket is on the map, push to stack
		if _, ok := m[bracket]; ok {
			stack = append(stack, bracket)
		} else {
			// if stack is empty, it means the parentheses is invalid
			// e.g. ]} or []} then return false
			if len(stack) == 0 {
				return false
			}

			// checks last index stack should be paired with current bracket
			// if not, return false
			// if paired, remove the last index
			lastIdx := len(stack) - 1
			if m[stack[lastIdx]] != bracket {
				return false
			} else {
				stack = append(stack[:lastIdx], stack[lastIdx+1:]...)
			}
		}
	}

	// if stack empty means parentheses is valid
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({{{{}}}))"))
}

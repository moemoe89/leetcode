package main

import "fmt"

func fizzBuzz(n int) []string {
	// create output variable
	var output []string

	// loop until n number
	for i := 1; i <= n; i++ {
		// default value
		answer := ""

		if i%15 == 0 { // if divisible by 3 & 5, means 3*5 = 15
			answer = "FizzBuzz"
		} else if i%3 == 0 { // if divisible by 3
			answer = "Fizz"
		} else if i%5 == 0 { // if divisible by 5
			answer = "Buzz"
		} else { // none above conditions are true
			answer = fmt.Sprintf("%d", i)
		}

		output = append(output, answer)
	}

	return output
}

func main() {
	fmt.Println(fizzBuzz(15))
}

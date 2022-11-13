package main

import "fmt"

func romanToInt(s string) int {
	// creates a map to store the symbol includes special symbol (IV, IX, etc)
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	// output integer
	output := 0
	// number of s index
	nS := len(s) - 1

	// skip the symbol if fall for special symbol (IV, IX, etc)
	skip := false

	// iterates s symbols
	for i, _ := range s {
		// skip if the value is true
		if skip {
			skip = false
			continue
		}

		// find current symbol (from backward)
		// e.g. MCMXCIV -> start from V, I, C, etc
		symbol := string(s[nS-i])

		// to find the special symbol, checks
		// if the next symbol index is exist
		nextSymbolIndex := nS - i - 1
		if nextSymbolIndex >= 0 {
			// if special symbol find, then use that special symbol
			// and do skip for the next symbol calculation
			// e.g. current index is V, the next iterate is I
			// then use symbol IV instead I + V
			beforeSymbol := string(s[nextSymbolIndex])
			if _, ok := m[beforeSymbol+symbol]; ok {
				skip = true
				symbol = beforeSymbol + symbol
			}
		}

		// sum all symbol integer value
		output += m[symbol]
	}

	// return the integer
	return output
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

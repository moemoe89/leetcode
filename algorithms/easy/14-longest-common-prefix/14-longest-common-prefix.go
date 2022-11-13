package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	// find the min character of each str
	// assign the first array of strs because it always >= 1
	minChar := len(strs[0])

	for i := range strs {
		if len(strs[i]) < minChar {
			minChar = len(strs[i])
		}
	}

	output := ""
	// count the number of strs
	nStrs := len(strs)

	// iterates char by char based on min char
	// e.g. ["flower","flow","flight"]
	// will check first char for each str
	// 1 -> (f,f,f), 2 -> (l,l,l), 3 -> (o,o,i)
	// then it will stop at 2 char (fl)
	for i := 0; i < minChar; i++ {
		// assign the first char and from first array of strs
		// because it always >= 1 and the if it's empty array,
		// should not go to this iteration
		char := strs[0][i]
		for j := 0; j < nStrs; j++ {
			if char != strs[j][i] {
				// if the char from other str is not same
				// return to exit block
				goto Exit
			}
		}

		// push the character to output variable
		output += string(char)
	}

	// return block
Exit:
	return output
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

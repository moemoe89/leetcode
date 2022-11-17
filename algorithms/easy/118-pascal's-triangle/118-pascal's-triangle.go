package main

import "fmt"

func generate(numRows int) [][]int {
	// create a return variable
	var rows [][]int

	// loop the row as the num rows input
	for i := 0; i < numRows; i++ {
		// create a row variable
		var row []int

		// loop the row as the num rows input
		for j := 0; j <= i; j++ {
			// default num value
			num := 1

			// if j position is not first and last
			// get the sum number of previous row, before and same index
			// for the example input is 3
			// first row because doesn't have previous row then become [1]
			// second row has previous row, but the index is not sufficient then become [1, 1]
			// third row has previous row, but only 2nd index has sufficient in the previous row,
			// which is the previous row before and same index is 1st and 2nd index, 1+1 = 2 then become [1, 2, 1]
			if j > 0 && j < i {
				num = rows[i-1][j-1] + rows[i-1][j]
			}

			row = append(row, num)
		}

		rows = append(rows, row)
	}

	return rows
}

func main() {
	fmt.Println(generate(3))
}

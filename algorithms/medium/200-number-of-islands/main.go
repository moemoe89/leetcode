package main

import "fmt"

func numIslands(grid [][]byte) int {
	island := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				island++

				dfs(grid, i, j)
			}
		}
	}

	return island
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))
}

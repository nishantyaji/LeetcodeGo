package main

import "fmt"

func countServers(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	total := 0
	rowFlags := make([]int, rows, rows)
	colFlags := make([]int, cols, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				total++
				rowFlags[i]++
				colFlags[j]++
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && rowFlags[i] == 1 && colFlags[j] == 1 {
				total--
			}
		}
	}
	return total
}

func main() {
	grid := [][]int{{1, 0}, {1, 1}}
	fmt.Println(countServers(grid))
}

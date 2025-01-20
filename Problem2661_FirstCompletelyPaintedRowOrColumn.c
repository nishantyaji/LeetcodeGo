package main

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])

	rowCount := make([]int, rows, rows)
	for i := 0; i < rows; i++ {
		rowCount[i] = cols
	}
	colCount := make([]int, cols, cols)
	for i := 0; i < cols; i++ {
		colCount[i] = rows
	}
	rowLabel := make([]int, len(arr), len(arr))
	colLabel := make([]int, len(arr), len(arr))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := mat[i][j]
			rowLabel[val-1] = i
			colLabel[val-1] = j
		}
	}
	for i, a := range arr {
		rowCount[rowLabel[a-1]]--
		colCount[colLabel[a-1]]--
		if rowCount[rowLabel[a-1]] == 0 || colCount[colLabel[a-1]] == 0 {
			return i
		}
	}
	return rows*cols - 1
}

func main() {
	fmt.Println(firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))
}

package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	blockLen := len(blocks)
	window := 0
	res := 0
	for i := 0; i < blockLen; i++ {
		temp := 0
		if blocks[i] == 'B' {
			temp = 1
		}
		if i >= k && blocks[i-k] == 'B' {
			temp -= 1
		}
		window += temp
		if window > res {
			res = window
		}
	}
	return k - res
}
func main() {
	fmt.Println("Hello, 世界")
}

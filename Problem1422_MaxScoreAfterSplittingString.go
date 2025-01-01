package main

import "fmt"

func maxScore(s string) int {
	res := -501
	zeroes := 0
	ones := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroes++
		} else {
			ones++
		}
		if i < len(s)-1 {
			res = max(res, zeroes-ones)
		}
	}
	return res + ones
}

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
	fmt.Println(maxScore("1111"))
	fmt.Println(maxScore("00"))
}

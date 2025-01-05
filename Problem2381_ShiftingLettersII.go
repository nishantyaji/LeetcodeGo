package main

import (
	"fmt"
	"strings"
)

func shiftingLetters(s string, shifts [][]int) string {
	dp := make([]int, len(s)+1, len(s)+1)

	for _, shift := range shifts {
		if shift[2] > 0 {
			dp[shift[0]]++
			dp[shift[1]+1]--
		} else {
			dp[shift[0]]--
			dp[shift[1]+1]++
		}
	}
	cum := 0
	aInt := int('a')
	zInt := int('z')
	var builder strings.Builder

	for i := 0; i < len(s); i++ {
		cum += dp[i]
		temp := int(s[i]) + ((cum%26)+26)%26
		if temp > zInt {
			temp = temp - zInt + aInt - 1
		}
		builder.WriteString(string(rune(temp)))
	}
	return builder.String()

}

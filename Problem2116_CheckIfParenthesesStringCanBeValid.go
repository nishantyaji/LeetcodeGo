package main

import "fmt"

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	opc := 0
	clc := 0

	for i, c := range locked {
		if c == '1' {
			if rune(s[i]) == '(' {
				opc++
			} else {
				clc++
			}
		}
		if opc > n/2 || clc > (i+1)/2 {
			return false
		}
	}
	opc = 0
	clc = 0
	for i := n - 1; i >= 0; i-- {
		if rune(locked[i]) == '1' {
			if rune(s[i]) == '(' {
				opc++
			} else {
				clc++
			}
		}
		if clc > n/2 || opc > (n-i)/2 {
			return false
		}
	}
	return true
}

func main() {
	s := "))()))"
	locked := "010100"
	fmt.Println(canBeValid(s, locked))
	s = "()()"
	locked = "0000"
	fmt.Println(canBeValid(s, locked))
	s = ")"
	locked = "0"
	fmt.Println(canBeValid(s, locked))
}

package main

import "fmt"

func minimumLength(s string) int {
	mp := make(map[rune]int)
	for _, c := range s {
		if _, ok := mp[c]; !ok {
			mp[c] = 0
		}
		mp[c]++
	}
	res := 0
	for _, v := range mp {
		if v%2 == 0 {
			res += 2
		} else {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(minimumLength("abaacbcbb"))
}

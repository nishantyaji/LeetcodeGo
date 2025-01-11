package main

import "fmt"

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; !ok {
			mp[s[i]] = 0
		}
		mp[s[i]]++
	}

	odds := 0
	for _, v := range mp {
		if v%2 == 1 {
			odds++
		}
	}
	if odds > k {
		return false
	}
	return true
}

func main() {
	fmt.Println(canConstruct("annabelle", 2))
}

package main

import (
	"fmt"
	"strings"
)

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if _, found := strings.CutPrefix(words[j], words[i]); found {
				iLen := len(words[i])
				jLen := len(words[j])
				if jLen < iLen {
					continue
				}
				suffix := 1
				for k := 0; k < iLen; k++ {
					if words[j][jLen-1-k] != words[i][iLen-1-k] {
						suffix = 0
						break
					}
				}
				if suffix == 1 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	test := []string{"a", "aba", "ababa", "aa"}
	res := countPrefixSuffixPairs(test)
	fmt.Println(res)
}

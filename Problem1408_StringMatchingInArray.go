package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func stringMatching(words []string) []string {
	slice := make([]string, 0, len(words))
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(words, lenCmp)
	flags := make([]int, len(words), len(words))
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				flags[i] = 1
				break
			}
		}
	}
	for i := 0; i < len(words); i++ {
		if flags[i] == 1 {
			slice = append(slice, words[i])
		}
	}
	return slice
}

func main() {
	test := []string{"as", "hero", "mass", "superhero"}
	res := stringMatching(test)
	for _, r := range res {
		fmt.Println(r)
	}
}

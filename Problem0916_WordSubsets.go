package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {
	uniMap := make(map[rune]int)
	res := make([]string, 0, len(words1))

	for _, w := range words2 {
		thisMap := make(map[rune]int)
		for _, c := range w {
			if _, ok := thisMap[c]; !ok {
				thisMap[c] = 0
			}
			thisMap[c] = thisMap[c] + 1
		}
		for k, v := range thisMap {
			if _, ok := uniMap[k]; !ok {
				uniMap[k] = v
			}
			uniMap[k] = max(v, uniMap[k])
		}
	}
	for _, w := range words1 {
		thisMap := make(map[rune]int)
		for _, c := range w {
			if _, ok := thisMap[c]; !ok {
				thisMap[c] = 0
			}
			thisMap[c] = thisMap[c] + 1
		}
		skip := false
		for k, v := range uniMap {
			if _, ok := thisMap[k]; !ok {
				skip = true
				break
			}
			if thisMap[k] < v {
				skip = true
				break
			}
		}
		if !skip {
			res = append(res, w)
		}
	}
	return res
}

func main() {
	w1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	w2 := []string{"e", "o"}
	res := wordSubsets(w1, w2)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

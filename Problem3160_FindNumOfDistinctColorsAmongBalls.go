package main

import "fmt"

func change(mp map[int]int, latest map[int]int, index int, val int) int {

	if prevVal, ok := latest[index]; ok {
		mp[prevVal] = mp[prevVal] - 1
		if mp[prevVal] == 0 {
			delete(mp, prevVal)
		}
	}

	if _, ok := mp[val]; !ok {
		mp[val] = 0
	}
	mp[val] += 1
	latest[index] = val
	return len(mp)
}

func queryResults(limit int, queries [][]int) []int {
	mp := make(map[int]int)
	latest := make(map[int]int)

	qLen := len(queries)
	res := make([]int, qLen, qLen)
	for i, q := range queries {
		res[i] = change(mp, latest, q[0], q[1])
	}
	return res
}

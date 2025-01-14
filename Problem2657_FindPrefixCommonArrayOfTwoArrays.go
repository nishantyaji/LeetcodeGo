package main

import "fmt"

func addMap(mp map[int]int, key int) {
	if _, ok := mp[key]; !ok {
		mp[key] = 0
	}
	mp[key]++
}

func findThePrefixCommonArray(A []int, B []int) []int {
	aMap := make(map[int]int)
	bMap := make(map[int]int)

	n := len(A)
	res := make([]int, len(A), len(A))
	res[len(A)-1] = len(A)
	for i := n - 1; i > 0; i-- {
		_, oka := bMap[A[i]]
		_, okb := aMap[B[i]]

		if oka && okb {
			res[i-1] = res[i]
		} else if A[i] == B[i] || oka || okb {
			res[i-1] = res[i] - 1
		} else {
			res[i-1] = res[i] - 2
		}

		addMap(aMap, A[i])
		addMap(bMap, B[i])
	}
	return res
}

func main() {
	p1 := []int{1, 3, 2, 4}
	p2 := []int{3, 1, 2, 4}
	res := findThePrefixCommonArray(p1, p2) // {0, 2, 3, 4}
	fmt.Println(res)
}

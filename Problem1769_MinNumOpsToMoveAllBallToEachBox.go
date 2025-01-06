package main

import "fmt"

func minOperations(boxes string) []int {
	sLen := len(boxes)
	length := 0
	var indices []int
	total := 0
	for i := 0; i < sLen; i++ {
		if boxes[i] == '1' {
			length++
			indices = append(indices, i)
			total += i
		}
	}

	res := make([]int, sLen, sLen)
	runIndex := 0
	right := length

	if length > 0 && indices[runIndex] == 0 {
		runIndex++
		right--
	}
	res[0] = total
	for i := 1; i < sLen; i++ {
		thisVal := res[i-1]
		exclude := 0

		if runIndex < length && indices[runIndex] == i {
			runIndex++
			thisVal--
			right--
			exclude = 1
		}

		thisVal = thisVal - right + (length - right - exclude)
		res[i] = thisVal
	}
	return res
}

func main() {
	res := minOperations("001011")
	for _, r := range res {
		fmt.Println(r)
	}
}

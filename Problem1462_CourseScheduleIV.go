package main

import "fmt"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
 	var dp = make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		dp[i] = make([]bool, numCourses, numCourses)
	}   
	
	for _, arr := range prerequisites {
		dp[arr[0]][arr[1]] = true
	}
	
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				dp[i][j] = dp[i][j] || dp[i][k] && dp[k][j]
			}
		}
	}
	resLen := len(queries)
	res := make([]int, resLen, resLen)
	for i, arr := range queries{
		res[i] := dp[arr[0]][arr[1]]
	}
	return res
}

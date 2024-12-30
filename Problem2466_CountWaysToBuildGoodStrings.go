package main

import "fmt"

func countGoodStrings(low int, high int, zero int, one int) int {
	mn := min(zero, one)
	mx := max(zero, one)
	slice := make([]int, mx+high+1, mx+high+1)
	slice[mn] += 1
	slice[mx] += 1
	base := 1000000007
	for i := mn; i <= high; i++ {
		if slice[i] > 0 {
			slice[i+zero] = (slice[i+zero] + slice[i]) % base
			slice[i+one] = (slice[i+one] + slice[i]) % base
		}
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + slice[i]) % base
	}
	return res

}

func main() {
	fmt.Println(countGoodStrings(2, 3, 1, 2))
	fmt.Println(countGoodStrings(3, 3, 1, 1))
}

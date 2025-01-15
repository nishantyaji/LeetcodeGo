package main

import (
	"fmt"
	"math"
	"math/bits"
)

func minimizeXor(num1 int, num2 int) int {
	len1 := bits.Len(uint(num1))
	ones2 := bits.OnesCount(uint(num2))
	base := int(math.Pow(2, float64(len1-1)))
	newBase := int(math.Pow(2, float64(len1)))
	i := 0
	ans := 0
	stack := make([]int, 0, len1)
	for i < ones2 && base >= 1 {
		if base&num1 > 0 {
			ans |= base
			i++
		} else {
			stack = append(stack, base)
		}
		base >>= 1
	}
	idx := len(stack)
	for i < ones2 {
		if idx > 0 {
			ans |= stack[idx-1]
			idx--
		} else {
			ans |= newBase
			newBase <<= 1
		}
		i++
	}
	return ans

}

func main() {
	fmt.Println(minimizeXor(3, 5))
	fmt.Println(minimizeXor(1, 12))
}

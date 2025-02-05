package main

import "fmt"

type CharCounter [26]int

func (cc *CharCounter) init(s string) {
	for _, r := range s {
		cc.add(r)
	}
}

func (cc *CharCounter) add(r rune) {
	index := int(r - 'a')
	cc[index]++
}

func (cc *CharCounter) remove(r rune) {
	index := int(r - 'a')
	cc[index]--
}

func (cc *CharCounter) isEqual(other CharCounter) bool {
	for i := 0; i < 26; i++ {
		if cc[i] != other[i] {
			return false
		}
	}
	return true
}

func areAlmostEqual(s1 string, s2 string) bool {
	c1 := &CharCounter{}
	c1.init(s1)
	c2 := &CharCounter{}
	c2.init(s2)

	if c1.isEqual(*c2) == false {
		return false
	}

	cnt := 0
	lens1 := len(s1)
	for i := 0; i < lens1; i++ {
		if s1[i] != s2[i] {
			cnt++
		}
		if cnt > 2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
	fmt.Println(areAlmostEqual("caa", "aab"))

}

package main

import "fmt"

type UnionFind struct {
	UfSize int
	Parent []int
	Size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		UfSize: n,
		Parent: make([]int, n),
		Size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.Parent[i] = i
		uf.Size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Union(i int, j int) {
	irep := uf.Find(i)
	jrep := uf.Find(j)
	if irep == jrep {
		return
	}

	isize := uf.Size[i]
	jsize := uf.Size[j]

	if isize < jsize {
		uf.Parent[irep] = jrep
		uf.Size[jrep] += uf.Size[irep]
	} else {
		uf.Parent[jrep] = irep
		uf.Size[irep] += uf.Size[jrep]
	}
}

func (uf UnionFind) Find(i int) int {
	if uf.Parent[i] != i {
		uf.Parent[i] = uf.Find(uf.Parent[i])
	}
	return uf.Parent[i]
}

func findRedundantConnection(edges [][]int) []int {
    mp := make(map[int]int)   
    res := []int{-1, -1}

    for _, edge := range edges {
        mp[edge[0]] = 1
        mp[edge[1]] = 1
    }
    n := len(mp)
    mp = make(map[int]int)
    uf := NewUnionFind(n)
    for _, edge := range edges {
        flag := true
        if _, ok1 := mp[edge[0]]; ok1 {
            if _, ok2 := mp[edge[1]]; ok2 {
                if uf.Find(edge[0] - 1) == uf.Find(edge[1] - 1) {
                    res[0] = edge[0]
                    res[1] = edge[1]
                    flag = false
                }
            }
        }
        if flag {
            mp[edge[0]] = 1
            mp[edge[1]] = 1
            uf.Union(edge[0] - 1, edge[1] - 1)
        }
    }
    return res
}

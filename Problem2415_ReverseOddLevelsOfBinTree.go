package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type deque []int

func (dq *deque) Push(i int) {
	*dq = append(*dq, i)
}

func (dq *deque) Pop() int {
	res := (*dq)[len(*dq)-1]
	*dq = (*dq)[:len(*dq)-1]
	return res
}

func (dq *deque) PopLeft() int {
	res := (*dq)[0]
	*dq = (*dq)[1:]
	return res
}

func (dq *deque) Len() int {
	return len(*dq)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	mp := make(map[int]*deque) // Note this

	if root != nil {
		recurse(root, mp, 0)
		assign(root, mp, 0)
	}
	return root
}

func insertMap(mp map[int]*deque, depth int, val int) {
	_, ok := mp[depth]
	if !ok {
		mp[depth] = &deque{} // Note this
	}
	mp[depth].Push(val)
}

func recurse(node *TreeNode, mp map[int]*deque, depth int) {
	insertMap(mp, depth, node.Val)
	if node.Left != nil {
		recurse(node.Left, mp, depth+1)
	}
	if node.Right != nil {
		recurse(node.Right, mp, depth+1)
	}
}

func assign(node *TreeNode, mp map[int]*deque, depth int) {
	if depth%2 == 1 {
		node.Val = mp[depth].Pop()
	} else {
		node.Val = mp[depth].PopLeft()
	}

	if node.Left != nil {
		assign(node.Left, mp, depth+1)
	}
	if node.Right != nil {
		assign(node.Right, mp, depth+1)
	}
}

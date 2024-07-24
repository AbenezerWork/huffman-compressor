package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Id    uint64
	Par   uint64
	Count int64
	Chr   byte
	Left  uint64
	Right uint64
}

type NodeList []Node

func (n NodeList) Len() int           { return len(n) }
func (n NodeList) Less(i, j int) bool { return n[i].Count < n[j].Count }
func (n NodeList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n *NodeList) Pop() any {
	old := *n
	h := len(old)
	x := old[h-1]
	*n = old[0 : h-1]
	return x
}

func (n *NodeList) Push(node any) {
	*n = append(*n, node.(Node))
}

type pair struct {
	node uint64
	pref []bool
}

var nodes map[uint64]Node

func BuildTable(root *Node) *map[byte][]bool {
	stk := []pair{}
	stk = append(stk, pair{node: root.Id, pref: []bool{}})
	table := make(map[byte][]bool)

	for len(stk) > 0 {
		curr := stk[len(stk)-1]
		pref := curr.pref
		currNode := nodes[curr.node]
		stk = stk[0 : len(stk)-1]
		if currNode.Chr != 0 {
			table[currNode.Chr] = pref
		} else {
			if currNode.Left != 0 {
				dup := make([]bool, len(pref))
				copy(dup, pref)
				dup = append(dup, false)
				stk = append(stk, pair{node: currNode.Left, pref: dup})

			}
			if currNode.Right != 0 {
				dup := make([]bool, len(pref))
				copy(dup, pref)
				dup = append(dup, true)
				stk = append(stk, pair{node: currNode.Right, pref: dup})

			}

		}
	}
	for key, value := range table {
		fmt.Println(rune(byte(key)), ": ", value)
	}
	return &table
}

var root Node

func buildTree() *map[byte][]bool {
	nodes = map[uint64]Node{}
	BuildHeap := &NodeList{}
	heap.Init(BuildHeap)
	var id uint64
	id++
	for key, value := range Cnt {
		node := Node{Id: id, Count: value, Chr: key}
		heap.Push(BuildHeap, node)
		nodes[id] = node
		id++
	}

	for BuildHeap.Len() > 1 {

		l := heap.Pop(BuildHeap).(Node)
		r := heap.Pop(BuildHeap).(Node)
		mix := Node{
			Id:    id,
			Count: l.Count + r.Count,
			Chr:   0,
			Left:  l.Id,
			Right: r.Id,
		}
		nodes[id] = mix
		id++
		heap.Push(BuildHeap, mix)
	}

	root = heap.Pop(BuildHeap).(Node)
	nodes[0] = root
	return BuildTable(&root)

}

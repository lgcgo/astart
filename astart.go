package main

import (
	"container/heap"
	"fmt"
)

type Noder interface {
	Neighbor() []Noder        // 周边查找器
	CalNeighborG(c *node) int // G值计算器
	CalNeighborH(e *node) int // H值计算器
}

type node struct {
	Noder Noder
	F     int
	G     int
	H     int
}

type nodes []*node

type nodeMap map[Noder]*node

type nodeLink map[Noder]Noder

// 实现按F值排序栈
func (ns nodes) Len() int           { return len(ns) }
func (ns nodes) Less(i, j int) bool { return ns[i].F < ns[j].F }
func (ns nodes) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns *nodes) Push(x interface{}) {
	*ns = append(*ns, x.(*node))
}
func (ns *nodes) Pop() interface{} {
	old := *ns
	n := len(old)
	x := old[n-1]
	*ns = old[0 : n-1]
	return x
}

func Astart(s, e Noder) []Noder {

	openList := make(nodeMap, 0)   // 开放列表
	closedList := make(nodeMap, 0) // 关闭列表
	linkList := make(nodeLink, 0)  // 关系列表

	sNode := &node{s, 0, 0, 0}
	eNode := &node{e, 0, 0, 0}

	oh := make(nodes, 0) // 开放节点的排序切片
	heap.Init(&oh)
	heap.Push(&oh, sNode) // 添加起点到排序列表

	openList[s] = sNode

	b := false // 找到了，直接跳出

	for {
		cNode := heap.Pop(&oh).(*node)  // 获取当前节点
		delete(openList, cNode.Noder)   // 移出开放列表
		closedList[cNode.Noder] = cNode // 加入关闭列表
		nb := cNode.Noder.Neighbor()
		for _, v := range nb {
			if _, ok := closedList[v]; ok {
				continue
			}
			if _, ok := openList[v]; ok {
				parentNoder := linkList[v]
				parentNode := closedList[parentNoder] // 因为v被考察过，所以其父节点必定存在关闭列表中
				if v.CalNeighborG(parentNode) < v.CalNeighborG(cNode) {
					cNode = parentNode
					heap.Push(&oh, cNode) // 重新放入排序中
					g := v.CalNeighborG(cNode)
					h := v.CalNeighborH(eNode)
					f := g + h
					node := &node{v, f, g, h}
					openList[v] = node
					break
				}
			}
			g := v.CalNeighborG(cNode)
			h := v.CalNeighborH(eNode)
			f := g + h
			node := &node{v, f, g, h}
			heap.Push(&oh, node)
			linkList[v] = cNode.Noder
			if v == e {
				b = true
				fmt.Println("Has find!")
				break
			}
			openList[v] = node
		}

		if len(openList) == 0 || b {
			if !b {
				fmt.Println("It is no find!")
			}
			fmt.Println("It is finish!")
			break
		}
	}

	res := make([]Noder, 0)
	if b {
		res = append(res, e)
		for {
			father, ok := linkList[e]
			res = append(res, father)
			if father == s || !ok {
				break
			}
			e = father
		}
	}

	return res
}

package main

import (
	"fmt"
	"log"
)

type Area struct {
	CntX  int              // X轴格子数量
	CntY  int              // Y轴格子数量
	Grids map[string]*Grid // 所有格子集合
}

type Grid struct {
	Area   *Area
	X      int
	Y      int
	IsLock bool
}

func (g *Grid) Neighbor() []Noder {
	x := g.X
	y := g.Y
	als := [4][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}}
	res := make([]Noder, 0)
	for _, v := range als {
		ag, ok := g.Area.Grids[calAlias(v[0], v[1])]
		if !ok || ag.IsLock { // 过滤锁定位置
			continue
		}
		res = append(res, ag)
	}
	return res
}

func (g *Grid) CalNeighborG(c *node) int {
	return c.G + 1
}

func (g *Grid) CalNeighborH(e *node) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	eg := e.Noder.(*Grid)
	return abs(eg.X-g.X) + abs(eg.Y-g.Y)
}

// 锁定
func (a *Area) LockGrid(x int, y int) bool {
	alias := calAlias(x, y)
	if _, ok := a.Grids[alias]; !ok {
		return false
	}
	a.Grids[alias].IsLock = true
	return true
}

// 获取
func (a *Area) GetGrid(x int, y int) *Grid {
	if _, ok := a.Grids[calAlias(x, y)]; !ok {
		log.Fatalf("Grid no find %s", calAlias(x, y))
	}
	return a.Grids[calAlias(x, y)]
}

// 计算格子ID
func calAlias(x int, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
}

// 创建
func CreateArea(CntX int, CntY int) *Area {
	a := &Area{
		CntX:  CntX,
		CntY:  CntY,
		Grids: make(map[string]*Grid),
	}
	for i := 0; i < CntX; i++ {
		for j := 0; j < CntY; j++ {
			a.Grids[calAlias(i, j)] = &Grid{
				Area:   a,
				X:      i,
				Y:      j,
				IsLock: false,
			}
		}
	}
	return a
}

// 打印地图
func (a *Area) P(c *Grid, e *Grid) {

	r := Astart(c, e)

	link := make(map[string]bool, 0)
	for _, v := range r {
		link[calAlias(v.(*Grid).X, v.(*Grid).Y)] = true
	}

	for i := 0; i < a.CntY; i++ {
		for j := 0; j < a.CntX; j++ {
			alias := calAlias(j, i)
			g := a.Grids[alias]
			_, ok := link[alias]
			if g.IsLock {
				fmt.Print("= ")
			} else if ok {
				fmt.Print("* ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}

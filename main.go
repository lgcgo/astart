package main

func main() {

	// 创建一个12*15的网格区域
	a := CreateArea(12, 15)

	// 任意设置路障
	a.LockGrid(0, 4)
	a.LockGrid(1, 4)
	a.LockGrid(2, 4)
	a.LockGrid(3, 4)
	a.LockGrid(4, 4)
	a.LockGrid(5, 4)
	a.LockGrid(6, 4)
	a.LockGrid(7, 4)
	a.LockGrid(8, 4)

	a.LockGrid(11, 7)
	a.LockGrid(10, 7)
	a.LockGrid(9, 7)
	a.LockGrid(8, 7)
	a.LockGrid(7, 7)
	a.LockGrid(6, 7)
	a.LockGrid(5, 7)
	a.LockGrid(4, 7)
	a.LockGrid(3, 7)

	a.LockGrid(0, 9)
	a.LockGrid(1, 9)
	a.LockGrid(2, 9)
	a.LockGrid(3, 9)
	a.LockGrid(4, 9)
	a.LockGrid(5, 9)
	a.LockGrid(6, 9)
	a.LockGrid(7, 9)
	a.LockGrid(8, 9)
	a.LockGrid(8, 10)
	a.LockGrid(8, 11)

	a.LockGrid(4, 14)
	a.LockGrid(4, 13)
	a.LockGrid(4, 12)

	c := a.GetGrid(1, 14) // 出发点current
	e := a.GetGrid(1, 0)  // 目的地end

	a.P(c, e) // 打印地图和路线

}

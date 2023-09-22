package models

type Coord struct {
	MaxX int // 点击边界最大 x
	MinX int // 点击边界最小 x
	MaxY int // 点击边界最大 y
	MinY int // 点击边界最小 y
	Cx   int // 中点坐标 x
	Cy   int // 中点坐标 y
}

func (c Coord) NewCoord() *Coord {
	return &c
}

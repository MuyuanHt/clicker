package work

import (
	"clicker/inits"
	"clicker/models"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

var (
	dx, dy int = 0, 0 // 最大最小坐标差值
	ex, ey int = 0, 0 // 随机生成目的坐标
)

// FirstLoc 初始化鼠标位置
func FirstLoc(c models.Coord) {
	robotgo.Move(c.Cx, c.Cy)
}

// RadiusMouse 设置鼠标偏移量 鼠标位置变化
func RadiusMouse(c models.Coord) {
	// 将当前鼠标移动到下一个坐标范围中心位置
	robotgo.Move(c.Cx, c.Cy)
	dx = c.MaxX - c.MinX // 计算差值生成随机位置
	dy = c.MaxY - c.MinY
	ex = rand.Intn(dx) + c.MinX //在差距范围内计算随机值+最小坐标得到范围内位置
	ey = rand.Intn(dy) + c.MinY
	robotgo.MoveSmooth(ex, ey, inits.Cfg.MouseSpeed, inits.Cfg.MouseASpeed) // 设置鼠标移动速度与加速度
}

// ChangeCoord 坐标变换 用于多个测定范围时使用
// m 为切片长度 n 为当前所在坐标范围在切片中的下标
func ChangeCoord(m, n int) int {
	if n == m-1 {
		return 0 // 当前位于切片末尾时回到切片前端
	}
	n++ // 移动到下一个区域
	return n
}

// OnceRunClick 每一轮的执行 包括坐标数量以内的点击与等待 num 表示每轮坐标变换次数 sub 表示下一个坐标
func OnceRunClick(cs []models.Coord, num, sub int) {
	for i := 0; i < num; i++ {
		sub = ChangeCoord(len(cs), sub)                                                    // 坐标变换 返回下一个指向的坐标范围
		RadiusMouse(cs[sub])                                                               // 控制坐标变换
		AutoClick()                                                                        // 控制点击事件
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(inits.Cfg.ChangeCoordTime))) // 1000ms 以内
	}
}

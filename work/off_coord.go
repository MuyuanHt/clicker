package work

import (
	"clicker/inits"
	"clicker/models"
	"github.com/go-vgo/robotgo"
	"log"
	"math/rand"
	"time"
)

var (
	x, y   int  = 0, 0  // 当前坐标
	dx, dy int  = 0, 0  // 最大最小坐标差值
	ex, ey int  = 0, 0  // 随机生成目的坐标
	is     bool = false // 判断是否为用户
)

// FirstLoc 初始化鼠标位置
func FirstLoc(c models.Coord) (int, int) {
	x = (c.MinX + c.MinX) / 2
	y = (c.MaxY + c.MinY) / 2
	robotgo.Move(x, y)
	return x, y
}

var stopCh = make(chan bool)

// RadiusMouse 设置鼠标偏移量 鼠标位置变化
func RadiusMouse(c models.Coord) {
	x, y = robotgo.GetMousePos() // 获取当前鼠标位置坐标
	dx = c.MaxX - c.MinX         // 计算差值生成随机位置
	dy = c.MaxY - c.MinY
	ex = rand.Intn(dx) + c.MinX //在差距范围内计算随机值+最小坐标得到范围内位置
	ey = rand.Intn(dy) + c.MinY
	go MoveMouse(c, ex, ey) // 在移动鼠标的同时判断是否为用户操作
	select {
	case is = <-stopCh:
		if is == true {
			useCh <- struct{}{}
		}
		return
	case <-time.After(time.Second * time.Duration(inits.Cfg.MaxOutTime)):
		log.Println("[超时] 将回到初始位置。")
		FirstLoc(c) // 回到中心点
	}
}

// MoveMouse 检测鼠标移动 同时判断是否为用户操作 ex ey 为目标位置
func MoveMouse(c models.Coord, ex, ey int) {
	x, y = robotgo.GetMousePos()
	if IsUser(x, y, c) == true {
		stopCh <- true
	} else {
		robotgo.MoveSmooth(ex, ey, inits.Cfg.MouseSpeed, inits.Cfg.MouseASpeed) // 设置鼠标移动速度与加速度
		stopCh <- false
	}
}

// IsUser 判断是否用户在操作鼠标 假设越界就是用户在使用鼠标
func IsUser(x, y int, c models.Coord) bool {
	if x > c.MaxX || x < c.MinX || y > c.MaxY || y < c.MinY {
		return true
	}
	return false
}

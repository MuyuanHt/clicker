package measure

import (
	"clicker/inits"
	"clicker/models"
	"clicker/utils"
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

// GetCoord 获取初始化鼠标范围坐标
// 该函数用于第一次运行时获取点击区域上下左右边界时使用
func GetCoord() {
	fmt.Println("准备开始测试 ")
	fmt.Printf("预期测试次数 %d 次\n", inits.Cfg.TestClickNum)
	fmt.Printf("测试准备时间 %d 秒\n", inits.Cfg.TestStartTime)
	fmt.Println("准备时间请将鼠标移动到需要点击的范围...")
	fmt.Println("点击鼠标左键提前结束准备时间...")
	tm := time.Duration(inits.Cfg.TestStartTime) * time.Second
	timer := time.NewTimer(tm)
	// 点击鼠标左键立即开始测定
	down := make(chan struct{})
	go func() {
		if robotgo.AddMouse("left") {
			down <- struct{}{}
		}
	}()
	select {
	case <-timer.C:
		fmt.Println("准备时间结束，测试开始...")
	case <-down:
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("选择立即测试，测试开始...")
	}
	i := 0 // 计数器 当达到测试次数时退出循环
	// 设定 max 为左上边界 min 为右下边界 刚好相反方便测试
	maxX, maxY, minX, minY := 0, 0, inits.Cfg.WindowMaxX, inits.Cfg.WindowMaxY
	x, y := 0, 0                                                   // 初始化鼠标位置
	tim := time.Duration(inits.Cfg.TestIntervalTime) * time.Second // 每次测试间隔时间
	for {
		i++
		time.Sleep(tim)
		x, y = robotgo.GetMousePos() // 获取当前鼠标位置坐标
		fmt.Printf("第 [ %d ] 次获得鼠标位置( %d , %d )\n", i, x, y)
		maxX = maxNum(maxX, x)
		maxY = maxNum(maxY, y)
		minX = minNum(minX, x)
		minY = minNum(minY, y)
		if i == inits.Cfg.TestClickNum {
			break
		}
	}
	fmt.Println("测试得到各边界值如下")
	fmt.Printf("MinX = %d, MaxX = %d\n", minX, maxX)
	fmt.Printf("MinY = %d, MaxY = %d\n", minY, maxY)
	coord := models.Coord{
		MaxX: maxX,
		MinX: minX,
		MaxY: maxY,
		MinY: minY,
		Cx:   (maxX + minX) / 2,
		Cy:   (maxY + minY) / 2,
	}
	utils.WriteCoord(coord)
}

func maxNum(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func minNum(m, n int) int {
	if m < n {
		return m
	}
	return n
}

package work

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

var (
	clickNum int = 0    // 同一时刻点击次数
	itm      int = 0    // 每次点击间隔时间
	x, y     int = 0, 0 // 鼠标当前位置
)

// AutoClick 控制鼠标点击事件
func AutoClick() {
	x, y = robotgo.GetMousePos() // 获取当前鼠标位置坐标
	fmt.Printf("鼠标点击位置 ( %d , %d )\n", x, y)
	clickNum = rand.Intn(5) + 1 // 设置连击次数范围为 [1,5]
	robotgo.Click("left", true, clickNum)
}

// IntervalTime 每次执行间隔时间
func IntervalTime(n int) {
	itm = rand.Intn(n) + 1 // 在 [1,max] 范围内随机间隔
	time.Sleep(time.Duration(itm) * time.Second)
}

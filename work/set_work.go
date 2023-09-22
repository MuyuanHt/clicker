package work

import (
	"bufio"
	"clicker/inits"
	"clicker/models"
	"clicker/utils"
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var down = make(chan struct{})
var useCh = make(chan struct{})
var memStats runtime.MemStats

// SetWork 设定执行时间与如何执行
func SetWork() {
	coord := utils.ReadCoord()   // 读取边界值
	FirstLoc(coord)              // 初始化鼠标位置
	finish := AllTime(inits.Cfg) // 设定程序运行时间
	timer := time.NewTimer(finish)
	// 点击鼠标右键提前结束程序
	go func() {
		if robotgo.AddMouse("right") {
			fmt.Println("[退出] 用户主动退出")
			os.Exit(1)
		}
	}()
	//判断鼠标位置超出范围过大时为用户操作
	go func() {
		for {
			<-useCh
			fmt.Println("[暂停] 用户正在操作鼠标")
			// 等待用户操作
			UserUse()
			is = false         // 清除标记
			FirstLoc(coord)    // 重新初始化
			down <- struct{}{} // 恢复 main 函数运行
		}
	}()
	for {
		select {
		case <-timer.C:
			return
		default:
			IntervalTime(inits.Cfg.IntervalTime) // 每次执行间隔时间
			AutoClick()                          // 控制点击
			RadiusMouse(coord)                   // 控制偏移
			x, y = robotgo.GetMousePos()
			// 用户使用鼠标则阻塞主线程
			// 被判定为用户使用之后 is = true 在下一个循环才会恢复 false
			if is == true {
				<-down
			}
			// 监测程序运行占用内存
			runtime.ReadMemStats(&memStats)
			//fmt.Println("当前程序占用的内存（字节）：", memStats.Alloc)
			if memStats.Alloc > uint64(inits.Cfg.MaxSize) {
				log.Fatal("[退出] 程序超过最大内存占用")
			}
		}
	}
}

// AllTime 计算程序预计运行总时间
func AllTime(cfg models.Config) time.Duration {
	h := time.Duration(cfg.AppTimeH) * time.Hour
	m := time.Duration(cfg.AppTimeM) * time.Minute
	s := time.Duration(cfg.AppTimeS) * time.Second
	return h + m + s
}

// UserUse 用户正在使用
// 设定用户在输入框中输入任何文本即可继续
func UserUse() {
	fmt.Println("输入任意字符继续执行 或 输入 exit 退出")
	reader := bufio.NewReader(os.Stdin)
	var str string
	var err error
	for {
		str, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("读取错误", err)
		}
		// 去除输入中的换行符
		str = strings.TrimSpace(str)
		if strings.EqualFold(str, "exit") {
			fmt.Println("[退出] 用户主动退出")
			os.Exit(1)
		}
		if str != "" {
			fmt.Println("[继续] 程序继续执行")
			break
		}
	}
}

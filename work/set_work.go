package work

import (
	"bufio"
	"clicker/inits"
	"clicker/models"
	"clicker/utils"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"runtime"
	"strings"
	"time"
)

var down = make(chan struct{}) // 用户操作时阻塞主协程
var memStats runtime.MemStats  // 测定实时内存占用
var is bool = false            // 判断是否为用户

// SetWork 设定执行时间与如何执行
func SetWork() {
	fmt.Println("准备开始点击")
	coords, csl := utils.ReadCoord() // 读取边界值与坐标切片长度
	FirstLoc(coords[0])              // 初始化鼠标位置
	finish := AllTime(inits.Cfg)     // 设定程序运行时间
	timer := time.NewTimer(finish)
	sub := 0 // 初始坐标为切片中第一个元素坐标
	// 点击鼠标右键暂停程序
	go func() {
		for {
			if robotgo.AddMouse("right") { // 监听事件时会阻塞协程
				is = true
				fmt.Println("[暂停] 用户需要操作鼠标")
				UserUse()
				is = false          // 清除标记
				FirstLoc(coords[0]) // 重新初始化
				down <- struct{}{}  // 恢复 main 函数运行
			}
		}
	}()
	for {
		select {
		case <-timer.C: // 控制超时 到达时间之后程序自动退出
			return
		default:
			IntervalTime(inits.Cfg.IntervalTime) // 每次执行间隔时间
			// 被判定为用户使用鼠标时 is = true 在下一个循环才会恢复 false
			if is == true {
				<-down // 阻塞协程
			}
			OnceRunClick(coords, csl, sub)
			// 监测程序运行占用内存
			runtime.ReadMemStats(&memStats)
			//fmt.Println("当前程序占用的内存（字节）：", memStats.Alloc)
			if memStats.Alloc > uint64(inits.Cfg.MaxSize) {
				inits.Logger.Println("[退出] 程序超过最大内存占用限制")
				close(down)
				os.Exit(1)
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
	fmt.Println("输入任意字符继续执行")
	fmt.Println("输入 exit 或 quit 退出")
	reader := bufio.NewReader(os.Stdin)
	var str string
	var err error
	for {
		str, err = reader.ReadString('\n')
		if err != nil {
			inits.Logger.Fatal("读取错误", err)
		}
		// 去除输入中的换行符
		str = strings.TrimSpace(str)
		if strings.EqualFold(str, "exit") || strings.EqualFold(str, "quit") {
			fmt.Println("[退出] 用户主动退出")
			close(down)
			os.Exit(1)
		}
		if str != "" {
			fmt.Println("[继续] 程序继续执行")
			break
		}
	}
}

package measure

import (
	"bufio"
	"clicker/inits"
	"clicker/models"
	"clicker/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// StartGet 多开测定 调用此函数可以进行多范围测定
func StartGet() {
	coords := make([]models.Coord, 0, 10)
	n := 1 // 默认为一个坐标范围
	fmt.Printf("输入测定坐标范围数目在 [1 , %d] 之间\n", inits.Cfg.MaxCoordNum)
	reader := bufio.NewReader(os.Stdin)
	var str string
	var err error
	for {
		str, err = reader.ReadString('\n')
		if err != nil {
			inits.Logger.Fatal("读取错误 ", err)
		}
		// 去除输入中的换行符
		str = strings.TrimSpace(str)
		n, err = strconv.Atoi(str)
		if err != nil {
			inits.Logger.Println("输入不符，转型错误 ", err)
		}
		if n <= inits.Cfg.MaxCoordNum && n > 0 {
			fmt.Printf("测定次数为 [ %d ]\n", n)
			break
		}
		fmt.Println("输入不符，请重新输入")
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("正在进行第 [ %d ] 次范围测定\n", i)
		coords = append(coords, GetCoord())                                // 将每次测定的数据追加到切片中
		time.Sleep(time.Second * time.Duration(inits.Cfg.TestBetweenTime)) // 各个范围测试时间间隔
	}
	utils.WriteCoord(coords)
}

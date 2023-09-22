package inits

import (
	"clicker/models"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	Cfg     models.Config
	LogFile *os.File
	Logger  *log.Logger
)

func init() {
	InitLogger()
	Cfg = ReadConfig()
}

// ReadConfig 载入配置文件
func ReadConfig() models.Config {
	rand.Seed(time.Now().UnixNano()) //初始化随机种子
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatal(err)
	}
	config := models.Config{
		AppTimeH: viper.GetInt("app.AppTimeH"),
		AppTimeM: viper.GetInt("app.AppTimeM"),
		AppTimeS: viper.GetInt("app.AppTimeS"),

		SaveCoordFile:   viper.GetString("app.SaveCoordFile"),
		IntervalTime:    viper.GetInt("app.IntervalTime"),
		ChangeCoordTime: viper.GetInt("app.ChangeCoordTime"),

		MaxSize:    viper.GetInt("app.MaxSize") * 1024 * 1024,
		MaxOutTime: viper.GetInt("app.MaxOutTime"),

		MaxCoordNum: viper.GetInt("app.MaxCoordNum"),

		TestClickNum:     viper.GetInt("test.TestClickNum"),
		TestStartTime:    viper.GetInt("test.TestStartTime"),
		TestIntervalTime: viper.GetInt("test.TestIntervalTime"),
		TestBetweenTime:  viper.GetInt("between.TestBetweenTime"),

		WindowMaxX: viper.GetInt("window.WindowMaxX"),
		WindowMaxY: viper.GetInt("window.WindowMaxY"),

		MouseSpeed:  viper.GetFloat64("mouse.MouseSpeed"),
		MouseASpeed: viper.GetFloat64("mouse.MouseASpeed"),
	}
	return config
}

// InitLogger 初始化日志配置
func InitLogger() {
	// 初始化日志文件
	var err error
	LogFile, err = os.OpenFile("app_log.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("文件日志错误") // 使用 Fatal 时强制退出程序
	}
	// Logger 自定义日志
	Logger = log.New(LogFile, "Clicker Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

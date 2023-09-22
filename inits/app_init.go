package inits

import (
	"clicker/models"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

var Cfg models.Config

func init() {
	Cfg = ReadConfig()
}

// ReadConfig 载入配置文件
func ReadConfig() models.Config {
	rand.Seed(time.Now().UnixNano()) //初始化随机种子
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	config := models.Config{
		AppTimeH: viper.GetInt("app.AppTimeH"),
		AppTimeM: viper.GetInt("app.AppTimeM"),
		AppTimeS: viper.GetInt("app.AppTimeS"),

		SaveCoordFile: viper.GetString("app.SaveCoordFile"),
		IntervalTime:  viper.GetInt("app.IntervalTime"),

		MaxSize:    viper.GetInt("app.MaxSize") * 1024 * 1024,
		MaxOutTime: viper.GetInt("app.MaxOutTime"),

		TestClickNum:     viper.GetInt("test.TestClickNum"),
		TestStartTime:    viper.GetInt("test.TestStartTime"),
		TestIntervalTime: viper.GetInt("test.TestIntervalTime"),

		WindowMaxX: viper.GetInt("window.WindowMaxX"),
		WindowMaxY: viper.GetInt("window.WindowMaxY"),

		MouseSpeed:  viper.GetFloat64("mouse.MouseSpeed"),
		MouseASpeed: viper.GetFloat64("mouse.MouseASpeed"),
	}
	return config
}

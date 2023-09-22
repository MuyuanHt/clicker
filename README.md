# Clicker

## 基本信息

Clicker 基于 RobotGo 实现，可以测定并记录显示器屏幕内的某一范围，并在这个范围内进行随机点击，同时 Clicker 使用 Go 1.17 进行编写。

Clicker 由测试器 measurer 与点击器 worker 两部分构成，measurer 用于测定需要进行点击的范围，worker 用于在点击范围内进行随机点击。

## 如何使用

**Clicker 的使用方式**

1. 进入想要使用的目录并从 GitHub 获取源代码

```bash
git clone https://github.com/MuyuanHt/clicker.git
```

2. 打开源代码并编译 measurer.go 与 worker.go 文件

```bash
go build measurer.go

go build worker.go
```

3. 将编译好的 measurer 与 worker 的可执行文件与 configs 目录和 coord.json 文件复制到想要使用的目录中（可选）

4. 执行编译后的 measurer 文件

5. 执行编译后的 worker 文件

## 注意事项

1. 在运行时应当首先执行 measurer 文件再执行 worker 文件，否则可能会由于未测定点击区域而出现异常
2. 可以在 configs 目录下对 config.yml 文件进行编辑以此对 Clicker 进行设置，以下是 config.yml 默认内容

```yml
app :
  AppTimeH : 1 # 程序累计运行时间 单位/h
  AppTimeM : 0 # 程序累计运行时间 单位/min
  AppTimeS : 0 # 程序累计运行时间 单位/s
  SaveCoordFile : coord.json # 测试数据记录文件名称
  IntervalTime : 7 # 每次点击最大间隔时间
  MaxSize : 30 # 程序运行最大占用内存 单位/M
  MaxOutTime : 10 # 最大超时时间

test :
  TestClickNum : 6 # 测定点击范围时鼠标点击次数
  TestStartTime : 5 # 测定前准备时间 单位/s
  TestIntervalTime : 2 # 每次测定间隔时间 单位/s

window :
  WindowMaxX : 1920 # 显示器最大宽度 x
  WindowMaxY : 1080 # 显示器最大高度 y

mouse :
  MouseSpeed : 100.0  # 鼠标移动速度
  MouseASpeed : 100.0 # 鼠标移动加速度
```

3. 在程序运行时可以根据输出的提示内容进行操作，如跳过、暂停、退出等
4. 请保证 configs 目录与 coord.json 文件、编译好的 measurer 文件、worker 文件位于同一目录下
5. 默认暂停：用户鼠标超出测定范围
6. 默认退出：用户主动点击鼠标右键

## 免责声明

1. Clicker 核心功能基于 RobotGo 实现，详情查看 RobotGo 官方地址：https://github.com/go-vgo/robotgo

2. Clicker 编写过程中，鉴于本人能力，代码存在很多不足，有许多值得修改的地方，还请原谅

3. Clicker 仅供学习参考

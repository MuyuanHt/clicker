package models

type Config struct {
	AppTimeH        int
	AppTimeM        int
	AppTimeS        int
	SaveCoordFile   string
	IntervalTime    int
	MaxSize         int
	MaxOutTime      int
	MaxCoordNum     int
	ChangeCoordTime int

	TestClickNum     int
	TestStartTime    int
	TestIntervalTime int
	TestBetweenTime  int

	WindowMaxX int
	WindowMaxY int

	MouseSpeed  float64
	MouseASpeed float64
}

func (c Config) NewConfig() *Config {
	return &c
}

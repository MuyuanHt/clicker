package utils

import (
	"clicker/inits"
	"clicker/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

// WriteCoord 将测试数据写入文件
func WriteCoord(cs []models.Coord) {
	file, err := os.OpenFile(inits.Cfg.SaveCoordFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		inits.Logger.Fatal("Could not open, error: ", err)
	}
	defer CloseFile(file)
	var data []byte
	data, err = json.Marshal(cs)
	if err != nil {
		inits.Logger.Println("Could not marshal, error: ", err)
	}
	err = ioutil.WriteFile(inits.Cfg.SaveCoordFile, data, 0644)
	if err != nil {
		inits.Logger.Panicln("Write err: ", err)
	}
}

// ReadCoord 从文件中读取测试数据 返回数据与切片长度
func ReadCoord() ([]models.Coord, int) {
	var coords []models.Coord
	file, err := os.OpenFile(inits.Cfg.SaveCoordFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		inits.Logger.Fatal("Could not open, error: ", err)
	}
	defer CloseFile(file)
	var data []byte
	data, err = ioutil.ReadFile(inits.Cfg.SaveCoordFile)
	if err != nil {
		inits.Logger.Panicln("Read err: ", err)
	}
	err = json.Unmarshal(data, &coords)
	if err != nil {
		inits.Logger.Println("Unmarshal err: ", err)
	}
	return coords, len(coords)
}

// CloseFile 用于关闭文件
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		inits.Logger.Println("Could not close, error:", err)
	}
}

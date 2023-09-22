package utils

import (
	"clicker/inits"
	"clicker/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// WriteCoord 将测试数据写入文件
func WriteCoord(c models.Coord) {
	file, err := os.OpenFile(inits.Cfg.SaveCoordFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Could not open, error: ", err)
	}
	defer CloseFile(file)
	var data []byte
	data, err = json.Marshal(c)
	if err != nil {
		log.Println("Could not marshal, error: ", err)
	}
	err = ioutil.WriteFile(inits.Cfg.SaveCoordFile, data, 0644)
	if err != nil {
		log.Panicln("Write err", err)
	}
}

// ReadCoord 从文件中读取测试数据
func ReadCoord() models.Coord {
	coord := models.Coord{}
	file, err := os.OpenFile(inits.Cfg.SaveCoordFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Could not open, error: ", err)
	}
	defer CloseFile(file)
	var data []byte
	data, err = ioutil.ReadFile(inits.Cfg.SaveCoordFile)
	if err != nil {
		log.Panicln("Read err", err)
	}
	err = json.Unmarshal(data, &coord)
	if err != nil {
		log.Println("Unmarshal err", err)
	}
	return coord
}

// CloseFile 用于关闭文件
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println("Could not close, error: ", err)
	}
}

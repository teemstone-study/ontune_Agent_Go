package main

import (
	"encoding/json"
	"os"
)

type SettingAgent struct {
	filePath        string
	fileName        string
	KafkaServerAddr string
	KafkaServerPort string
	HeaderKeyCode   string
	StartNum        int
	EndNum          int
	GetPerfDataTime PerfInterval
}
type PerfInterval struct {
	Basic    int
	Process  int
	IO       int
	CoreUtil int
}

func (asetting *SettingAgent) Save() {
	jsonEncode, err := json.Marshal(asetting)
	if err != nil {
		panic(err)
	} else {
		nfile, err := os.Create(asetting.GetSavePath())
		if err != nil {
			panic(err)
		}
		nfile.Write(jsonEncode)
		defer nfile.Close()
	}
}

func (asetting *SettingAgent) Load() {
	savePath := asetting.GetSavePath()
	_, err := os.Stat(savePath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(asetting.filePath, os.ModeAppend)
			if err != nil {
				panic(err)
			}
			file, err := os.Create(savePath)
			if err != nil {
				panic(err)
			}

			jsonEncode, err := json.Marshal(asetting)
			if err != nil {
				panic(err)
			}
			file.Write(jsonEncode)
			defer file.Close()
		}
	}

	ontuneSettingData, err := os.Open(savePath)
	defer ontuneSettingData.Close()
	jsonDecoder := json.NewDecoder(ontuneSettingData)
	jsonDecoder.Decode(&asetting)
	asetting.Save()
}

func (asetting *SettingAgent) GetSavePath() string {
	return asetting.filePath + asetting.fileName
}

func initSetting(basePath string, filename string) *SettingAgent {
	asetting := SettingAgent{}
	asetting.filePath = basePath
	asetting.fileName = filename
	asetting.KafkaServerAddr = "127.0.0.1"
	asetting.KafkaServerPort = "9092"
	asetting.HeaderKeyCode = "testAgent"
	asetting.StartNum = 0
	asetting.EndNum = 100
	asetting.GetPerfDataTime = PerfInterval{}
	asetting.GetPerfDataTime.Basic = 1
	asetting.GetPerfDataTime.Process = 5
	asetting.GetPerfDataTime.IO = 5
	asetting.GetPerfDataTime.CoreUtil = 5

	return &asetting
}

func NewSetting(basePath string, filename string) *SettingAgent {
	asetting := initSetting(basePath, filename)
	asetting.Load()
	return asetting
}

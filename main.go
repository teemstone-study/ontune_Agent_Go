package main

import (
	"encoding/json"
	"fmt"
	"ontune_Kafka_Controller_Go/onTuneKafkaController"
	"ontune_Kafka_DataStruct/kafkaDataStruct"
	"sync"
	"time"
)

const (
	INIFILENAME string = "agent.json"
	INIFILEPATH string = `config\`
	FILEVERSION string = "0.0.0.1"
)

var wg sync.WaitGroup
var exitControl chan bool

func testhostSender(settingValue *SettingAgent) {
	exitControl = make(chan bool, 1)
	var hostList []kafkaDataStruct.HostAgentInfo
	hostList = make([]kafkaDataStruct.HostAgentInfo, settingValue.EndNum-settingValue.StartNum)
	for i := settingValue.StartNum; i < settingValue.EndNum; i++ {
		hostinfo := getNewHost(settingValue.HeaderKeyCode, i)
		hostList[i] = *hostinfo

		hostEncodeData, err := json.Marshal(hostinfo)
		if err != nil {
			fmt.Println(err)
		} else {
			onTuneKafkaController.SendKafkaData("host", hostinfo.AgentID, hostEncodeData)
			if settingValue.LOGMODE {
				fmt.Println("Sendhost:", hostinfo.AgentID)
			}
		}
	}

	basicPerfTick := time.NewTicker(time.Second * time.Duration(settingValue.GetPerfDataTime.Basic))
	basicPIDTick := time.NewTicker(time.Second * time.Duration(settingValue.GetPerfDataTime.CoreUtil))
	basicDiskNetTick := time.NewTicker(time.Second * time.Duration(settingValue.GetPerfDataTime.IO))

	for {
		select {
		case _ = <-basicPerfTick.C:
			{
				for _, hostinfo := range hostList {
					nowTime := time.Now().UTC()
					realtimePerf := getMakeOntuneRealTimePerfData(hostinfo, nowTime)
					hostEncodeData, err := json.Marshal(realtimePerf)
					if err != nil {
						fmt.Println(err)
					} else {
						onTuneKafkaController.SendKafkaData("realtimeperf", hostinfo.AgentID, hostEncodeData)
						if settingValue.LOGMODE {
							fmt.Println("SendRealTimePerf:", hostinfo.AgentID, " ", realtimePerf.Agenttime)
						}
					}
				}
			}

		case _ = <-basicPIDTick.C:
			{
				for _, hostinfo := range hostList {
					nowTime := time.Now().UTC()
					realtimePIDPerf := getMakeOntuneRealTimePIDData(hostinfo, nowTime)
					hostEncodeData, err := json.Marshal(realtimePIDPerf)
					if err != nil {
						fmt.Println(err)
					} else {
						onTuneKafkaController.SendKafkaData("realtimeppid", hostinfo.AgentID, hostEncodeData)
						if settingValue.LOGMODE {
							fmt.Println("SendRealTimePID:", hostinfo.AgentID, " ", realtimePIDPerf.Agenttime)
						}
					}
				}
			}
		case _ = <-basicDiskNetTick.C:
			{
				for _, hostinfo := range hostList {
					nowTime := time.Now().UTC()
					realtimeDisk := getMakeOntuneRealTimeDiskData(hostinfo, nowTime)
					realTimeDiskEncodeData, err := json.Marshal(realtimeDisk)
					if err != nil {
						fmt.Println(err)
					} else {
						onTuneKafkaController.SendKafkaData("realtimedisk", hostinfo.AgentID, realTimeDiskEncodeData)
						if settingValue.LOGMODE {
							fmt.Println("SendRealTimeDISK:", hostinfo.AgentID, " ", realtimeDisk.Agenttime)
						}
					}

					realtimeNet := getMakeOntuneRealTimeNetData(hostinfo, nowTime)
					realTimeNetEncodeData, err := json.Marshal(realtimeNet)
					if err != nil {
						fmt.Println(err)
					} else {
						onTuneKafkaController.SendKafkaData("realtimenet", hostinfo.AgentID, realTimeNetEncodeData)
						if settingValue.LOGMODE {
							fmt.Println("SendRealTimeNET:", hostinfo.AgentID, " ", realtimeNet.Agenttime)
						}
					}
				}
			}
		case <-exitControl:
			{
				close(exitControl)
				wg.Done()
				break
			}
		}
	}
}

func main() {
	settingValue := NewSetting(INIFILEPATH, INIFILENAME)
	kafkaconfig := onTuneKafkaController.SettingKafka{KafkaServerAddr: settingValue.KafkaServerAddr, KafkaServerPort: settingValue.KafkaServerPort}
	onTuneKafkaController.KafkaProducerControllerInit(&kafkaconfig)

	wg.Add(1)
	//일단 여기서 테스트
	go testhostSender(settingValue)

	wg.Wait()

}

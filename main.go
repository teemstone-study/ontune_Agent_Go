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
	for i := 0; i < settingValue.EndNum-settingValue.StartNum; i++ {
		hostinfo := getNewHost(settingValue.HeaderKeyCode, settingValue.StartNum+i)
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
				fmt.Println("SendPerfCount", len(hostList))
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
						onTuneKafkaController.SendKafkaData("realtimepid", hostinfo.AgentID, hostEncodeData)
						if settingValue.LOGMODE {
							fmt.Println("SendRealTimePID:", hostinfo.AgentID, " ", realtimePIDPerf.Agenttime)
						}
					}
				}
				fmt.Println("SendPIDCount", len(hostList))
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
				fmt.Println("SendNet,Disk Count", len(hostList))
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
	fmt.Println("Open Setting File")
	settingValue := NewSetting(INIFILEPATH, INIFILENAME)
	fmt.Println("Open Setting File Complete")
	fmt.Println("HeaderName : ", settingValue.HeaderKeyCode)
	fmt.Println("AgentStartNum : ", settingValue.StartNum)
	fmt.Println("AgentEndNum : ", settingValue.EndNum)
	fmt.Println("AgentCount : ", settingValue.EndNum-settingValue.StartNum)
	fmt.Println("Kafka Setting")
	kafkaconfig := onTuneKafkaController.SettingKafka{KafkaServerAddr: settingValue.KafkaServerAddr, KafkaServerPort: settingValue.KafkaServerPort}
	onTuneKafkaController.KafkaProducerControllerInit(&kafkaconfig)
	fmt.Println("Kafka Setting Complete (", *onTuneKafkaController.ProducerInit, ")")
	if *onTuneKafkaController.ProducerInit == true {
		fmt.Println("Agent Start")
		wg.Add(1)
		go testhostSender(settingValue)
		wg.Wait()
		fmt.Println("Agent Stop")
	}
	fmt.Println("Agent Close")
}

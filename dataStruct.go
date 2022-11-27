package main

import (
	"encoding/json"
	"os"
	"time"
)

//필요 설정 항목
///내 정보
//
///성능 정보

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

func NetSetting(basePath string, filename string) *SettingAgent {
	asetting := initSetting(basePath, filename)
	asetting.Load()
	return asetting
}

type HostAgentInfo struct {
	AgentName    string
	AgentID      string
	Model        string
	Serial       string
	Ip           string
	Os           string
	Agentversion string
	ProcessCount int
	ProcessClock int
	MemorySize   int
	SwapMemory   int
}

type RealTimePerf struct {
	AgentID          string
	Agenttime        time.Time
	User             int
	Sys              int
	Wait             int
	Idle             int
	ProcessorCount   int
	RunQueue         int
	BlockQueue       int
	WaitQueue        int
	PQueue           int
	PCRateUser       int
	PCRateSys        int
	MemorySize       int
	MemoryUsed       int
	MemoryPinned     int
	MemorySys        int
	MemoryUser       int
	MemoryCache      int
	Avm              int
	PagingspaceIn    int
	PaingSpaceOut    int
	FileSystemIn     int
	FileSystmeOut    int
	MemoryScan       int
	MemoryFreed      int
	SwapSize         int
	SwapUsed         int
	SwapActive       int
	Fork             int
	Exec             int
	Interupt         int
	SystemCall       int
	ContextSwitch    int
	Semaphore        int
	Msg              int
	DiskReadWrite    int
	DiskIOPS         int
	NetworkReadWrite int
	TopCommandID     int
	TopCommandCount  int
	TopUserID        int
	TopCPU           int
	TopDiskID        int
	TopvgID          int
	TOPBusy          int
	MaxPID           int
	ThreadCount      int
	PIDCount         int
	LinuxBuffer      int
	LinuxCached      int
	Linuxsrec        int
	Memused_mb       int
	IRQ              int
	SoftIRQ          int
	Swapused_MB      int
	DUSM             int
}

type RealTimePID struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []RealTimePIDInner
}

type RealTimePIDInner struct {
	Pid        int
	Ppid       int
	Uid        int
	Cmdname    string
	Username   string
	Argname    string
	Usr        int
	Sys        int
	Usrsys     int
	Sz         int
	Rss        int
	Vmem       int
	Chario     int
	Processcnt int
	Threadcnt  int
	Handlecnt  int
	Stime      int
	Pvbytes    int
	Pgpool     int
}

type RealTimeDisk struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []RealTimeDiskInner
}

type RealTimeDiskInner struct {
	Ioname       string
	Readrate     int
	Writerate    int
	Iops         int
	Busy         int
	Descname     string
	Readsvctime  int
	Writesvctime int
}

type RealTimeNet struct {
	AgentID   string
	Agenttime time.Time
	PerfList  []RealTimeNetInner
}

type RealTimeNetInner struct {
	Ioname    string
	Readrate  int
	Writerate int
	Readiops  int
	Writeiops int
	Errorps   int
	Collision int
}

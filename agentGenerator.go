package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getNewHost(headString string, number int) *HostAgentInfo {
	ahostData := HostAgentInfo{}
	ahostData.AgentID = headString + strconv.Itoa(number)
	ahostData.AgentName = headString + strconv.Itoa(number) + strconv.Itoa(rand.Intn(10000))
	ahostData.Agentversion = FILEVERSION
	ahostData.Ip = fmt.Sprintf("192.168.%d.%d", rand.Intn(250), rand.Intn(250))
	ahostData.MemorySize = rand.Intn(1000000000)
	ahostData.Model = "noModel"
	ahostData.Os = "Linux"
	ahostData.ProcessClock = rand.Intn(5000)
	ahostData.ProcessCount = rand.Intn(32)
	ahostData.Serial = strconv.Itoa(rand.Intn(100000000))
	ahostData.SwapMemory = rand.Intn(1000000000)

	return &ahostData
}

func getMakeOntuneRealTimePerfData(hostItem HostAgentInfo, baseTime time.Time) (arealTimeperf RealTimePerf) {
	arealTimeperf.AgentID = hostItem.AgentID
	arealTimeperf.Agenttime = baseTime
	arealTimeperf.User = rand.Intn(100)
	arealTimeperf.Sys = rand.Intn(100)
	arealTimeperf.Wait = rand.Intn(100)
	arealTimeperf.Idle = rand.Intn(100)
	arealTimeperf.ProcessorCount = rand.Intn(100)
	arealTimeperf.RunQueue = rand.Intn(100)
	arealTimeperf.BlockQueue = rand.Intn(100)
	arealTimeperf.WaitQueue = rand.Intn(100)
	arealTimeperf.PQueue = rand.Intn(100)
	arealTimeperf.PCRateUser = rand.Intn(100)
	arealTimeperf.PCRateSys = rand.Intn(100)
	arealTimeperf.MemorySize = rand.Intn(1000000000)
	arealTimeperf.MemoryUsed = rand.Intn(1000000000)
	arealTimeperf.MemoryPinned = rand.Intn(1000000000)
	arealTimeperf.MemorySys = rand.Intn(1000000000)
	arealTimeperf.MemoryUser = rand.Intn(1000000000)
	arealTimeperf.MemoryCache = rand.Intn(1000000000)
	arealTimeperf.Avm = rand.Intn(100)
	arealTimeperf.PagingspaceIn = rand.Intn(100)
	arealTimeperf.PaingSpaceOut = rand.Intn(100)
	arealTimeperf.MemoryScan = rand.Intn(100)
	arealTimeperf.MemoryFreed = rand.Intn(100)
	arealTimeperf.SwapSize = rand.Intn(100)
	arealTimeperf.SwapUsed = rand.Intn(100)
	arealTimeperf.SwapActive = rand.Intn(100)
	arealTimeperf.Fork = rand.Intn(100)
	arealTimeperf.Exec = rand.Intn(100)
	arealTimeperf.Interupt = rand.Intn(100)
	arealTimeperf.SystemCall = rand.Intn(100)
	arealTimeperf.ContextSwitch = rand.Intn(100)
	arealTimeperf.Semaphore = rand.Intn(100)
	arealTimeperf.Msg = rand.Intn(100)
	arealTimeperf.DiskReadWrite = rand.Intn(100)
	arealTimeperf.DiskIOPS = rand.Intn(100)
	arealTimeperf.NetworkReadWrite = rand.Intn(100)
	arealTimeperf.DiskIOPS = rand.Intn(100)
	arealTimeperf.NetworkReadWrite = rand.Intn(100)
	arealTimeperf.TopCommandID = rand.Intn(100)
	arealTimeperf.TopCommandCount = rand.Intn(100)
	arealTimeperf.TopUserID = rand.Intn(100)
	arealTimeperf.TopCPU = rand.Intn(100)
	arealTimeperf.TopDiskID = rand.Intn(100)
	arealTimeperf.TopvgID = rand.Intn(100)
	arealTimeperf.TOPBusy = rand.Intn(100)
	arealTimeperf.MaxPID = rand.Intn(100)
	arealTimeperf.ThreadCount = rand.Intn(100)
	arealTimeperf.PIDCount = rand.Intn(100)
	arealTimeperf.LinuxBuffer = rand.Intn(100)
	arealTimeperf.LinuxCached = rand.Intn(100)
	arealTimeperf.Linuxsrec = rand.Intn(100)
	arealTimeperf.Memused_mb = rand.Intn(100)
	arealTimeperf.IRQ = rand.Intn(100)
	arealTimeperf.SoftIRQ = rand.Intn(100)
	arealTimeperf.Swapused_MB = rand.Intn(100)
	arealTimeperf.DUSM = rand.Intn(100)

	return arealTimeperf

}

func getMakeOntuneRealTimePIDData(hostItem HostAgentInfo, baseTime time.Time) (arealTimecpu RealTimePID) {

	arealTimecpu.AgentID = hostItem.AgentID
	arealTimecpu.Agenttime = baseTime

	pidcount := rand.Intn(100)
	arealTimecpu.PerfList = make([]RealTimePIDInner, pidcount)

	for i := 0; i < pidcount; i++ {
		arealTimepidInner := RealTimePIDInner{}

		arealTimepidInner.Pid = rand.Intn(100)
		arealTimepidInner.Ppid = rand.Intn(100)
		arealTimepidInner.Uid = rand.Intn(100)
		var cmdname string
		switch rand.Intn(3) {
		case 1:
			cmdname = "java"
		case 2:
			cmdname = "node"
		case 3:
			cmdname = "ontuned"
		}
		arealTimepidInner.Cmdname = cmdname

		var username string
		switch rand.Intn(4) {
		case 1:
			username = "root"
		case 2:
			username = "gdm"
		case 3:
			username = "telegraf"
		case 4:
			username = "kibana"
		}
		arealTimepidInner.Username = username

		var argname string
		switch rand.Intn(4) {
		case 1:
			argname = "root"
		case 2:
			argname = "gdm"
		case 3:
			argname = "telegraf"
		case 4:
			argname = "kibana"
		}
		arealTimepidInner.Argname = argname

		arealTimepidInner.Usr = rand.Intn(100)
		arealTimepidInner.Sys = rand.Intn(100)
		arealTimepidInner.Usrsys = rand.Intn(100)
		arealTimepidInner.Sz = rand.Intn(100)
		arealTimepidInner.Rss = rand.Intn(100)
		arealTimepidInner.Vmem = rand.Intn(100)
		arealTimepidInner.Chario = rand.Intn(100)
		arealTimepidInner.Processcnt = rand.Intn(100)
		arealTimepidInner.Threadcnt = rand.Intn(100)
		arealTimepidInner.Handlecnt = rand.Intn(100)
		arealTimepidInner.Stime = rand.Intn(100)
		arealTimepidInner.Pvbytes = rand.Intn(100)

		arealTimecpu.PerfList[i] = arealTimepidInner
	}
	return arealTimecpu

}

func getMakeOntuneRealTimeDiskData(hostItem HostAgentInfo, baseTime time.Time) (arealTimeDisk RealTimeDisk) {

	arealTimeDisk.AgentID = hostItem.AgentID
	arealTimeDisk.Agenttime = baseTime

	diskcount := rand.Intn(20)
	arealTimeDisk.PerfList = make([]RealTimeDiskInner, diskcount)

	for i := 0; i < diskcount; i++ {
		arealTimeDiskInner := RealTimeDiskInner{}

		var diskioname string
		switch rand.Intn(4) {
		case 1:
			diskioname = "C:"
		case 2:
			diskioname = "D:"
		case 3:
			diskioname = "E:"
		case 4:
			diskioname = "F:"
		}
		arealTimeDiskInner.Ioname = diskioname
		arealTimeDiskInner.Readrate = rand.Intn(100)
		arealTimeDiskInner.Writerate = rand.Intn(100)
		arealTimeDiskInner.Iops = rand.Intn(100)
		arealTimeDiskInner.Busy = rand.Intn(100)

		var descname string
		switch rand.Intn(4) {
		case 1:
			descname = "centos_centos7kvm"
		case 2:
			descname = "centos_numatest"
		case 3:
			descname = "ontubun1604-vg"
		case 4:
			descname = "2a-ce75126bbade"
		}
		arealTimeDiskInner.Descname = descname
		arealTimeDiskInner.Readsvctime = rand.Intn(100)
		arealTimeDiskInner.Writesvctime = rand.Intn(100)

		arealTimeDisk.PerfList[i] = arealTimeDiskInner
	}

	return arealTimeDisk

}

func getMakeOntuneRealTimeNetData(hostItem HostAgentInfo, baseTime time.Time) (arealTimeNet RealTimeNet) {
	arealTimeNet.AgentID = hostItem.AgentID
	arealTimeNet.Agenttime = baseTime

	netcount := rand.Intn(5)
	arealTimeNet.PerfList = make([]RealTimeNetInner, netcount)

	for i := 0; i < netcount; i++ {
		arealTimenetInner := RealTimeNetInner{}
		var netioname string
		switch rand.Intn(4) {
		case 1:
			netioname = "Intel[R] 82574L Gigabit Network Connection"
		case 2:
			netioname = "Intel[R] PRO_1000 MT Network Connection"
		case 3:
			netioname = "Intel[R] 82574L Gigabit Network Connection _2"
		case 4:
			netioname = "VG_XenStorage-628ef03e-1cf7-cd4"
		}
		arealTimenetInner.Ioname = netioname
		arealTimenetInner.Readrate = rand.Intn(100)
		arealTimenetInner.Writerate = rand.Intn(100)
		arealTimenetInner.Readiops = rand.Intn(100)
		arealTimenetInner.Writeiops = rand.Intn(100)
		arealTimenetInner.Errorps = rand.Intn(100)
		arealTimenetInner.Collision = rand.Intn(100)

		arealTimeNet.PerfList[i] = arealTimenetInner
	}
	return arealTimeNet

}

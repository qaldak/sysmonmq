package systeminfo

import (
	"os/exec"

	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	logger "github.com/qaldak/SysMonMQ/logging"
	"golang.org/x/sys/unix"
)

type memStats struct {
	Total, Used, Free, Available uint64
}

func GetSystemInfo() (string, error) {
	cpuvags := getCPULoadStats()
	logger.Info(cpuvags)

	cpuTemp := getCPUTemp()
	logger.Info(cpuTemp)

	memoryStats := getMemoryStats()
	logger.Info(memoryStats)

	diskStats := getDiskStats()
	logger.Info(diskStats)

	uptimeStats := getUptimeStats()
	logger.Info(uptimeStats)

	lastStats := getLastLoginStats()
	logger.Info(lastStats)

	return "Foo", nil
}

func getCPULoadStats() *loadavg.Stats {
	cpuavgs, err := loadavg.Get()
	if err != nil {
		logger.Error("Failed to get cpu load avarage")
	}

	return cpuavgs
}

func getCPUTemp() float64 {
	// Todo: determine directyl from os
	return 49.2
}

func getDiskStats() uint64 {
	// Todo: determine directly from os
	diskStats, err := disk.Get()
	if err != nil {
		logger.Error("Faild to get disk stats")
	}

	logger.Info(diskStats)

	return 12
}

func getLastLoginStats() (string) {
	cmd := exec.Command("w", "-h", "-i")
	last, err := cmd.Output()
	if err != nil {
		logger.Error("Failed to get login informations")
	}
	logger.Info(string(last))
	return "Bar"
}

func getMemoryStats() *memStats {
	memoryStats, err := memory.Get()
	if err != nil {
		logger.Error("Failed to get memory stats")
	}

	// convert memStats in Mb and map
	mem := &memStats{
		Total:     memoryStats.Total / 1024 / 1024,
		Used:      memoryStats.Used / 1024 / 1024,
		Free:      memoryStats.Free / 1024 / 1024,
		Available: memoryStats.Available / 1024 / 1024,
	}

	return mem
}

func getUptimeStats() int64 {
	var info unix.Sysinfo_t
	if err := unix.Sysinfo(&info); err != nil {
		logger.Error(err)
	}
	logger.Info("Uptime ", info.Uptime)
	return info.Uptime
}



package systeminfo

import (
	"math"

	logger "github.com/qaldak/sysmonmq/internal/logging"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type memStats struct {
	Total, Used, Free, Available uint64
	UsedPercent                  float64
}

type diskStats struct {
	Path              string
	Total, Used, Free uint64
	UsedPercent       float64
}

// Get system informations from host
func GetSystemInfo() (*SystemInfo, error) {
	cpuavgs := getCPULoadStats()

	cpuTemp := getCPUTemp()

	memoryStats := getMemoryStats()

	diskStats := getDiskStats()

	uptimeStats := getUptimeStats()

	systemInfo := &SystemInfo{
		CPU01:          cpuavgs.Load1,
		CPU05:          cpuavgs.Load5,
		CPU15:          cpuavgs.Load15,
		CPU_temp:       cpuTemp,
		RAM_total:      memoryStats.Total,
		RAM_free:       memoryStats.Free,
		RAM_avlbl:      memoryStats.Available,
		RAM_used:       memoryStats.Used,
		Disk_total:     diskStats.Total,
		Disk_free:      diskStats.Free,
		Disk_used:      diskStats.Used,
		Sys_Uptime:     uptimeStats,
		LastLogin_date: "",
		LastLogin_user: "",
		LastLogin_from: "",
	}

	return systemInfo, nil
}

// Determine CPU load average (1 min, 5 min, 15 min)
func getCPULoadStats() *load.AvgStat {
	cpuavgs, err := load.Avg()
	if err != nil {
		logger.Error("Failed to get cpu load avarage")
	}

	logger.Info("CPU load avg: ", cpuavgs)

	return cpuavgs
}

// Determine CPU temperature informations
func getCPUTemp() float64 {
	t := 0.00

	tempStat, err := host.SensorsTemperatures()
	if err != nil || tempStat == nil {
		logger.Error("Failed to get cpu informations")
	}

	if len(tempStat) > 0 {
		logger.Info("Foo ", tempStat)
		t = tempStat[0].Temperature
	}

	logger.Info("Temp", t)

	return t
}

// Determine disk usage for root directory ("/")
func getDiskStats() *diskStats {
	usageStats, err := disk.Usage("/")
	if err != nil {
		logger.Error("Faild to get disk stats")
	}

	diskUsage := &diskStats{
		Total:       usageStats.Total / (1 << 30),                 // convert to Gb
		Free:        usageStats.Free / (1 << 30),                  // convert to Gb
		Used:        usageStats.Used / (1 << 30),                  // convert to Gb
		UsedPercent: math.Round(usageStats.UsedPercent*100) / 100, // convert to 2 decimal
	}

	logger.Info("Disk usage: ", diskUsage)

	return diskUsage
}

// Determine memory usage
func getMemoryStats() *memStats {
	mem, err := mem.VirtualMemory()
	if err != nil {
		logger.Error("Failed to get memory stats")
	}

	// convert memStats in Mb and map
	memUsage := &memStats{
		Total:       mem.Total / (1 << 20),
		Used:        mem.Used / (1 << 20),
		Free:        mem.Free / (1 << 20),
		Available:   mem.Available / (1 << 20),
		UsedPercent: math.Round(mem.UsedPercent*100) / 100, // convert to 2 decimal
	}

	logger.Info("Memory usage: ", memUsage)

	return memUsage
}

// Determine uptime from host
func getUptimeStats() uint64 {
	uptime, err := host.Uptime()
	if err != nil {
		logger.Error("Failed to get uptime info")
	}

	logger.Info("System uptime: ", uptime)

	return uptime
}

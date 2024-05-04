package systeminfo

import (
	"math"

	logger "github.com/qaldak/sysmonmq/logging"
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

	return "Foo", nil
}

// Determine CPU load average (1 min, 5 min, 15 min)
func getCPULoadStats() *load.AvgStat {
	cpuavgs, err := load.Avg()
	if err != nil {
		logger.Error("Failed to get cpu load avarage")
	}

	return cpuavgs
}

// Determine CPU temperature informations
func getCPUTemp() float64 {
	// Todo: implement
	// head -n 1 /sys/class/thermal/thermal_zone0/temps | xargs -I{} awk "BEGIN {printf \"%.2f\n\", {}/1000}")

	cpu, err := host.Info()
	if err != nil {
		logger.Error("Failed to get cpu informations")
	}

	logger.Info("cpu:", cpu)

	return 0.00
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
		UsedPercent: math.Round(mem.UsedPercent*100) / 100,
	}

	logger.Info("memory: ", memUsage)

	return memUsage
}

// Determine uptime from host
func getUptimeStats() uint64 {
	uptime, err := host.Uptime()
	if err != nil {
		logger.Error("Failed to get uptime info")
	}
	logger.Info(uptime)

	return uptime
}

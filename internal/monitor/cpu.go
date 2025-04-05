package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu" // it is very useful for work with CPUs level and python psutil version for Golang
)

func GetCPUUsageBar() string {
	percent, _ := cpu.Percent(0, false)
	// 0 -- means measure CPU usage instantly no delay
	// false -- means get overall CPU usage not per core
	usage := percent[0]

	return fmt.Sprintf("%s  %.2f%%", GetBar(usage), usage)
}

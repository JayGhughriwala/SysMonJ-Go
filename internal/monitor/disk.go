package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsageBar() string {
	// here I select the root path "/" for windows edit this to "C:\"
	usage, _ := disk.Usage("/") // here you can select path as your choice
	return fmt.Sprintf("%s  %.2f%% (%.2f / %.2f GB)",
		GetBar(usage.UsedPercent),
		usage.UsedPercent,
		float64(usage.Used)/1e9,
		float64(usage.Total)/1e9)

}

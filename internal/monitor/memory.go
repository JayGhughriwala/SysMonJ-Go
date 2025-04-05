package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryUsageBar() string {
	vm, _ := mem.VirtualMemory()
	return fmt.Sprintf("%s   %.2f%% (%.2f / %.2f GB)", GetBar(vm.UsedPercent), vm.UsedPercent, float64(vm.Used)/1e9, float64(vm.Total)/1e9)
}

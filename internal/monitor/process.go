package monitor

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcInfo struct {
	PID    int32
	Name   string
	CPU    float64
	Memory float32 // here I returns float32 because below memorypercent returns 32byte size float so

}

func PrintTopProcesses(limit int) {
	procs, _ := process.Processes()
	var top []ProcInfo
	for _, p := range procs { // i dont care about index give me the process itself
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()

		if cpu > 0.5 {
			top = append(top, ProcInfo{PID: p.Pid, Name: name, CPU: cpu, Memory: mem})
		}
	}
	sort.Slice(top, func(i, j int) bool { // sorting which process have high cpu usage
		return top[i].CPU > top[j].CPU
	})
	fmt.Printf(" %-8s | %-17s | %-7s | %-7s\n", "PID", "Name", "CPU %", "MEM %")
	fmt.Println("----------|-------------------|--------|--------")

	for i, p := range top {
		if i >= limit {
			break
		}
		fmt.Printf(" %-8d | %-17s | %6.2f%% | %6.2f%%\n", p.PID, truncate(p.Name, 17), p.CPU, p.Memory)
	}

}

// this function generally helps in shorting the process name
func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}

	return s
}

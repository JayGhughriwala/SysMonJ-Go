package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

func GetNetworkStats() string {
	io, _ := net.IOCounters(false)
	stats := io[0]
	return fmt.Sprintf("↑ %.2f MB   ↓ %.2f MB", float64(stats.BytesSent)/1e6, float64(stats.BytesRecv)/1e6)
}

// ↑   ↓

package main

// importing some useful module for requirements
import (
	"fmt"
	"time"

	"SysMonJ-Go/internal/monitor" // this both are for local practice
	"SysMonJ-Go/internal/utils"   //

	//"github.com/JayGhughriwala/SysMonJ-Go/internal/monitor"
	//"github.com/JayGhughriwala/SysMonJ-Go/internal/utils"
	"github.com/fatih/color"
)

func main() {
	for {
		utils.ClearScreen()
		// here set blue color
		color.Blue("╔═══════════════════════════════════════════════════════╗")
		color.Blue("║               🌐  SysMonJ-Go Monitoring Tool          ║")
		color.Blue("╠═══════════════════════════════════════════════════════╣")
		color.Blue("║            Real-time System Resource Monitor          ║")
		color.Blue("║       Built with 💙 in Go | MIT Licensed | v1.0.0     ║")
		color.Blue("╚═══════════════════════════════════════════════════════╝\n")

		//cyan color set

		color.Cyan("📊 SYSTEM OVERVIEW")
		fmt.Println("─────────────────────────────────────────────────────────")
		fmt.Printf("🧠 CPU Usage      : %s\n", monitor.GetCPUUsageBar())
		fmt.Printf("🧠 Memory Usage   : %s\n", monitor.GetMemoryUsageBar())
		fmt.Printf("💾 Disk Usage     : %s\n", monitor.GetDiskUsageBar())
		fmt.Printf("🌐 Network        : %s\n\n", monitor.GetNetworkStats())

		// red color for process

		color.Red("🔥 TOP PROCESSES BY CPU")
		fmt.Println("─────────────────────────────────────────────────────────")
		monitor.PrintTopProcesses(5)
		// here I set 5 because I am only interested to view top five processes
		fmt.Println("\n💡 Tip: Press Ctrl+C to exit. Refreshes every 3 seconds.")
		time.Sleep(3 * time.Second)
		// here above line is useful for the set your own refreshing time
	}
}

// ╔ ═ ╗ ║ ╚ ╝  🌐 💙 💡 🔥 🌐 💾 🧠 📊

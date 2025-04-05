package monitor

import "strings"

func GetBar(percent float64) string {
	total := 20
	filled := int((percent / 100.0) * float64(total))
	return "[" + strings.Repeat("█", filled) + strings.Repeat("░", total-filled) + "]"

}

// █  ░

package monitor

import "strings"

func GetBar(percent float64) string {
	total := 20
	filler := int((percent / 100.0) * float64(total))
	return "[" + strings.Repea
}

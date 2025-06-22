package formatter

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Purple = "\033[35m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Cyan   = "\033[36m"
)

func PrintHeader(title string) {
	bar := strings.Repeat("═", len(title)+6)
	fmt.Printf("\n%s%s╔══ %s ══╗%s\n", Cyan, Bold, title, Reset)
	fmt.Printf("%s%s║  %s  ║%s\n", Cyan, Bold, title, Reset)
	fmt.Printf("%s%s╚═══%s═══╝%s\n", Cyan, Bold, bar, Reset)
}

func PrintValue(key string, value int) {
	var color string
	switch {
	case value > 50:
		color = Red
	case value > 25:
		color = Yellow
	case value > 10:
		color = Cyan
	default:
		color = Green
	}

	bar := strings.Repeat("█", value) 
	fmt.Printf("%s%-20s%s : %s%s%s\n", Purple, key, Reset, color, bar, Reset)
}

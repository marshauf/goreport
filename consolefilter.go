package goreport

import (
	"fmt"
)

const (
	ConsoleColorDefaultForeground = 39
	ConsoleColorBlack             = 30
	ConsoleColorRed               = 31
	ConsoleColorGreen             = 32
	ConsoleColorYellow            = 33
	ConsoleColorBlue              = 34
	ConsoleColorMagenta           = 35
	ConsoleColorCyan              = 36
	ConsoleColorLightGray         = 37
	ConsoleColorDarkGray          = 90
	ConsoleColorLightRed          = 91
	ConsoleColorLightGreen        = 92
	ConsoleColorLightYellow       = 93
	ConsoleColorLightBlue         = 94
	ConsoleColorLightMagenta      = 95
	ConsoleColorLightCyan         = 96
	ConsoleColorWhite             = 97
)

type ConsoleFilter struct {
	SeverityColor map[Severity]int
}

func NewConsoleFilter() Filter {
	return &ConsoleFilter{
		SeverityColor: map[Severity]int{
			Debug:   ConsoleColorDefaultForeground,
			Info:    ConsoleColorBlue,
			Warning: ConsoleColorYellow,
			Error:   ConsoleColorLightRed,
			Fatal:   ConsoleColorRed,
			Panic:   ConsoleColorLightMagenta,
		},
	}
}

func (f *ConsoleFilter) Filter(entry Entry) bool {
	for key, value := range entry {
		switch s := value.(type) {
		case Severity:
			if color, ok := f.SeverityColor[s]; ok {
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, entry[key])
			}
		}
	}
	return false
}

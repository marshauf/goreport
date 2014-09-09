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
}

func NewConsoleFilter() Filter {
	return &ConsoleFilter{}
}

func (f *ConsoleFilter) Filter(entry Entry) {
	for key, value := range entry {
		switch key {
		case EntryKeySeverity:
			switch value {
			case Debug:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorDefaultForeground, entry[key])
			case Info:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorBlue, entry[key])
			case Warning:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorYellow, entry[key])
			case Error:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorLightRed, entry[key])
			case Fatal:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorRed, entry[key])
			case Panic:
				entry[key] = fmt.Sprintf("\x1b[%dm%s\x1b[0m", ConsoleColorLightMagenta, entry[key])
			}
		case EntryKeyTime:
		case EntryKeyMessage:
		}
	}
}

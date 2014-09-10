package goreport

type SeverityFilter struct {
	Level int
}

func NewSeverityFilter(level int) Filter {
	return &SeverityFilter{
		Level: level,
	}
}

func (sf *SeverityFilter) Filter(entry Entry) {
	if value, ok := entry[EntryKeySeverity]; ok {
		switch level := value.(type) {
		case Severity:
			if level.Less(sf.Level) {
				entry.Clear()
			}
		}
	}
}

package goreport

import (
	"fmt"
	"io"
	"time"
)

type TextFormatter struct {
	TimeFormat     string
	EntrySeperator string
}

func NewTextFormatter() Formatter {
	return &TextFormatter{TimeFormat: time.RFC3339, EntrySeperator: "\n"}
}

func (f *TextFormatter) Write(entry Entry, w io.Writer) error {
	var err error

	// Has severity entry
	if value, ok := entry[EntryKeySeverity]; ok {
		switch severity := value.(type) {
		case string, error: // Is of type string or has function String() string
			_, err = fmt.Fprintf(w, "%s", severity)
		default:
			_, err = fmt.Fprintf(w, "%v", severity)
		}

		if err != nil {
			return err
		}
	}

	// Has time entry
	if value, ok := entry[EntryKeyTime]; ok {
		switch t := value.(type) {
		case time.Time:
			_, err = fmt.Fprintf(w, "(%s)", t.Format(f.TimeFormat))
		case string, error: // Is of type string or has function String() string
			_, err = fmt.Fprintf(w, "(%s)", t)
		case func() string:
			_, err = fmt.Fprintf(w, " %s=%s", EntryKeyTime, t())
		case func(Entry) string:
			_, err = fmt.Fprintf(w, " %s=%s", EntryKeyTime, t(entry))
		default:
			_, err = fmt.Fprintf(w, "(%v)", t)
		}
		if err != nil {
			return err
		}
	}

	// Has message entry
	if value, ok := entry[EntryKeyMessage]; ok {
		switch message := value.(type) {
		case string, error:
			_, err = fmt.Fprintf(w, " %s", message)
		case func() string:
			_, err = fmt.Fprintf(w, " %s=%s", EntryKeyMessage, message())
		case func(Entry) string:
			_, err = fmt.Fprintf(w, " %s=%s", EntryKeyMessage, message(entry))
		default:
			_, err = fmt.Fprintf(w, " %v", message)
		}
		if err != nil {
			return err
		}
	}

	// Write the rest of the entry
	for key, value := range entry {
		// Ignore all previously written entry data
		switch key {
		case EntryKeySeverity, EntryKeyTime, EntryKeyMessage:
			continue
		default:
			switch v := value.(type) {
			case func() string:
				_, err = fmt.Fprintf(w, " %s=%s", key, v())
			case func(Entry) string:
				_, err = fmt.Fprintf(w, " %s=%s", key, v(entry))
			default:
				_, err = fmt.Fprintf(w, " %s=%v", key, v)
			}
		}
	}
	_, err = fmt.Fprint(w, f.EntrySeperator)
	return err
}

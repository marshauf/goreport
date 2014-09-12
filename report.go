package goreport

import (
	"fmt"
	"io"
	"time"
)

const (
	LogBufferSize = 8
)

type Report struct {
	Filters []Filter
	entries []Entry
	n       int
}

func New() *Report {
	return &Report{
		entries: make([]Entry, LogBufferSize),
	}
}

func (r *Report) Debug(format string, args ...interface{}) {
	r.Log(Debug, format, args...)
}

func (r *Report) Info(format string, args ...interface{}) {
	r.Log(Info, format, args...)
}

func (r *Report) Warning(format string, args ...interface{}) {
	r.Log(Warning, format, args...)
}

func (r *Report) Error(format string, args ...interface{}) {
	r.Log(Error, format, args...)
}

func (r *Report) Fatal(format string, args ...interface{}) {
	r.Log(Fatal, format, args...)
}

func (r *Report) Panic(format string, args ...interface{}) {
	r.Log(Panic, format, args...)
}

func (r *Report) Log(s Severity, format string, args ...interface{}) {
	entry := make(Entry)
	entry[EntryKeySeverity] = s
	entry[EntryKeyTime] = time.Now()
	entry[EntryKeyMessage] = fmt.Sprintf(format, args...)
	r.Add(entry)
}

func (r *Report) Report(formatter Formatter, output io.Writer) {
	var err error
	for _, entry := range r.entries {
		if entry == nil {
			continue
		}
		entryCopy := entry.Copy()
		ignore := r.filter(entryCopy)
		if ignore || entryCopy == nil || len(entryCopy) == 0 {
			continue
		}
		err = formatter.Write(entryCopy, output)
		if err != nil {
			panic(err) // TODO panic with report.entries + error
		}
	}
}

func (r *Report) Add(entry Entry) {
	if r.n >= len(r.entries) {
		r.entries = append(r.entries, entry)
	} else {
		r.entries[r.n] = entry
	}
	r.n++
	return
}
func (r *Report) AddFilters(filters ...Filter) {
	if r.Filters == nil {
		r.Filters = filters
		return
	}
	r.Filters = append(r.Filters, filters...)
}

func (r *Report) filter(entry Entry) bool {
	for _, filter := range r.Filters {
		if filter == nil {
			continue
		}
		ignore := filter.Filter(entry)
		if ignore {
			return true
		}
	}
	return false
}

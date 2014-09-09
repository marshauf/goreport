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
		r.filter(entry)
		if entry == nil || len(entry) == 0 {
			continue
		}
		err = formatter.Write(entry, output)
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

// Clear removes all entries from the report
func (r *Report) Clear() {
	for i := range r.entries {
		r.entries[i].Clear()
	}
	r.n = 0
}

func (r *Report) AddFilters(filters ...Filter) {
	if r.Filters == nil {
		r.Filters = filters
		return
	}
	r.Filters = append(r.Filters, filters...)
}

func (r *Report) filter(entry Entry) {
	for _, filter := range r.Filters {
		if filter == nil {
			continue
		}
		filter.Filter(entry)
	}
}

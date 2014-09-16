goreport [![Build Status](https://travis-ci.org/marshauf/goreport.svg?branch=master)](https://travis-ci.org/marshauf/goreport)[![Coverage Status](https://coveralls.io/repos/marshauf/goreport/badge.png)](https://coveralls.io/r/marshauf/goreport)
========

A report library for go. Not stable!
Instead of logging everything and generating endless log files, goreport forms logs into reports.
At any point in a process the report can be exported.

# Features
+ Log collecting into reports
+ Middleware via filters (formatting, filtering, ....)
+ Multiple outputs
+ Output control, only output if report contains an error (reduces log size and only logs things which are important)

# Example

```go
package main

import (
  report "github.com/marshauf/goreport"
  "os"
)

func main() {
	r := report.New()
	r.Info("Simple log reports with %s.", "goreport")
	r.Add(report.Entry{
		"severity": Info,
		"features": []string{"middleware", "filters", "formatters", "optional output", "multiple output"},
		"simple": "unpolluted, formatted logs",
	})
	if r.HasError() {
		jf := &report.JsonFormatter{PrettyPrint:false}
		sendEmail := NewEmailSender()
		r.Report(jf, sendEmail)
		sendDB := NewDBWriter()
		if err := r.Report(jf, sendDB); err != nil {
			r.Fatal("Could not send email report: %s", err)
		}
	}
	r.AddFilters(NewConsoleFilter())
	if err := r.Report(NewTextFormatter(), os.Stdout); err != nil {
		panic(err) // Yes even reporting (errors) can fail
	}
}
```
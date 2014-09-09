package goreport

import (
	"bytes"
	"testing"
)

func TestReport(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Recovered from a panic: %v", r)
		}
	}()

	var (
		b  *bytes.Buffer = new(bytes.Buffer)
		b2 *bytes.Buffer = new(bytes.Buffer)
		f                = NewTextFormatter()
		r                = New()
	)
	if r == nil {
		t.Fatal("Report is nil")
	}
	if f == nil {
		t.Fatal("TextFormatter is nil")
	}

	// Try to uncommon, not intended inputs to r.Log
	r.Log(nil, "", nil)
	r.Log(nil, "", f.Write)

	r.Add(Entry{})
	r.Add(Entry{
		"severity": nil,
		"time":     nil,
		"message":  nil,
	})
	r.Add(Entry{
		"test": nil,
	})
	r.Add(Entry{
		"severity": "test",
		"time":     "unknown",
		"message":  func() string { return "message from a function" },
		"test":     0,
	})

	r.Log(Debug, "%s", &struct {
		Name string
	}{Name: "color test for debug"})
	r.Log(Info, "%s", &struct {
		Name string
	}{Name: "color test for info"})
	r.Log(Warning, "%s", &struct {
		Name string
	}{Name: "color test for warning"})
	r.Log(Error, "%s", &struct {
		Name string
	}{Name: "color test for error"})
	r.Log(Fatal, "%s", &struct {
		Name string
	}{Name: "color test for fatal"})
	r.Log(Panic, "%s", &struct {
		Name string
	}{Name: "color test for pannic"})

	r.AddFilters(NewConsoleFilter())

	r.Report(f, b)

	r.Clear()

	r.Report(f, b2)
	if b2.Len() > 0 {
		t.Errorf("After call to r.Clear, r.Report wrote to output: %s", b2.String())
	}
}

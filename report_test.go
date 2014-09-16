package goreport

import (
	"bytes"
	"testing"
)

func TestReport(t *testing.T) {
	var (
		err error
		b   *bytes.Buffer = new(bytes.Buffer)
		b2  *bytes.Buffer = new(bytes.Buffer)
		f                 = NewTextFormatter()
		r                 = New()
		r2                = New()
	)
	if r == nil {
		t.Fatal("Report is nil")
	}
	if f == nil {
		t.Fatal("TextFormatter is nil")
	}

	// Try uncommon, not intended inputs to r.Log
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

	r.AddFilters(NewConsoleFilter(), NewSeverityFilter(3))

	err = r.Report(f, b)
	if err != nil {
		t.Error(err)
	}
	err = r.Report(f, b2)
	if err != nil {
		t.Error(err)
	}

	if b.String() != b2.String() {
		t.Error("b != b2")
	}

	r2.Debug("%s", "debug")
	r2.Info("%s", "info")
	r2.Warning("%s", "warning")
	r2.Error("%s", "error")
	r2.Fatal("%s", "fatal")
	r2.Panic("%s", "panic")
	b.Reset()
	err = r2.Report(NewJsonFormatter(), b)
	if err != nil {
		t.Error(err)
	}
}

func TestReadmeExample(t *testing.T) {
	var err error

	r := New()
	r.Info("Simple log reports with %s.", "goreport")
	r.Add(Entry{
		"severity": Info,
		"features": []string{"middleware", "filters", "formatters", "optional output", "multiple output"},
		"simple":   "unpolluted, formatted logs",
	})
	if r.HasError() {
		jf := &JsonFormatter{PrettyPrint: false}
		sendEmail := NewEmailSender()
		r.Report(jf, sendEmail)
		sendDB := NewDBWriter()
		err = r.Report(jf, sendDB)
		if err != nil {
			t.Error(err)
		}
	}
	r.AddFilters(NewConsoleFilter())
	err = r.Report(NewTextFormatter(), &NullWriter{}) // os.Stdout
	if err != nil {
		t.Error(err)
	}
}

type NullWriter struct{}

func (w *NullWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func NewDBWriter() *NullWriter {
	return &NullWriter{}
}

func NewEmailSender() *NullWriter {
	return &NullWriter{}
}

package goreport

import (
	"encoding/json"
	"io"
)

// JsonFormatter
// Note: Only supports encoding/json types as Entry values
type JsonFormatter struct {
}

func NewJsonFormatter() Formatter {
	return &JsonFormatter{}
}

func (f *JsonFormatter) Write(entry Entry, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(entry)
}

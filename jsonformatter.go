package goreport

import (
	"encoding/json"
	"io"
)

// JsonFormatter
// Note: Only supports encoding/json types as Entry values
type JsonFormatter struct {
	PrettyPrint bool
	Indent      string
	Prefix      string
}

func NewJsonFormatter() Formatter {
	return &JsonFormatter{
		PrettyPrint: true,
		Indent:      "\t",
		Prefix:      "",
	}
}

func (f *JsonFormatter) Write(entry Entry, w io.Writer) error {
	if f.PrettyPrint {
		b, err := json.MarshalIndent(entry, f.Prefix, f.Indent)
		if err != nil {
			return err
		}
		_, err = w.Write(b)
		return err
	}

	e := json.NewEncoder(w)
	return e.Encode(entry)
}

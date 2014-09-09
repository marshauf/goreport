package goreport

import (
	"io"
)

type Formatter interface {
	Write(Entry, io.Writer) error
}

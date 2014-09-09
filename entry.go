package goreport

const (
	EntryKeySeverity = "severity"
	EntryKeyTime     = "time"
	EntryKeyMessage  = "message"
)

type Entry map[string]interface{}

func (e Entry) Clear() {
	for key := range e {
		delete(e, key)
	}
}

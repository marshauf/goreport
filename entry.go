package goreport

const (
	EntryKeySeverity = "severity"
	EntryKeyTime     = "time"
	EntryKeyMessage  = "message"
)

type Entry map[string]interface{}

func (e Entry) Copy() Entry {
	c := make(Entry)
	for k, v := range e {
		c[k] = v
	}
	return c
}

func (e Entry) Add(key string, value interface{}) Entry {
	e[key] = value
	return e
}
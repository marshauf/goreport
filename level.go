package goreport

var (
	Debug   = NewLevel("DEBUG", 0)
	Info    = NewLevel("INFO", 1)
	Warning = NewLevel("WARNING", 2)
	Error   = NewLevel("ERROR", 3)
	Fatal   = NewLevel("FATAL", 4)
	Panic   = NewLevel("PANIC", 5)
)

type Severity interface {
	String() string
	Severity() int
	Less(int) bool
}

// Level implements the Severity interface.
// It stores the severity value and the name of the severity.
// A higher severity value means more severe.
type Level struct {
	Name  string
	Value int
}

func NewLevel(name string, value int) *Level {
	return &Level{Name: name, Value: value}
}

func (l *Level) String() string {
	return l.Name
}

// Less reports if this level of severity is less severe than v.
func (l *Level) Less(v int) bool {
	return l.Value < v
}

func (l *Level) Severity() int {
	return l.Value
}

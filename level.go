package gol

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type Level int8

func (l Level) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}
	return ";?"
}

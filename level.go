package gol

import "fmt"

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type Level int8

func (l Level) String() (string, error) {
	switch l {
	case TRACE:
		return "TRACE", nil
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARN:
		return "WARN", nil
	case ERROR:
		return "ERROR", nil
	}
	return "", fmt.Errorf("Undefinted level: %s", l)
}

func Parse(level string) (Level, error) {
	switch level {
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	}
	return 0, fmt.Errorf("Unkown level: %s", level)
}

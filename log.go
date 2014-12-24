package gol

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	level  Level
	writer io.Writer
}

func ClassicLogger() *Logger {
	return &Logger{
		level:  INFO,
		writer: os.Stderr,
	}
}

func LevelLogger(level Level) *Logger {
	return &Logger{
		level:  level,
		writer: os.Stderr,
	}
}

func New(level Level, writer io.Writer) *Logger {
	return &Logger{
		level:  level,
		writer: writer,
	}
}

func (l *Logger) Level(level Level) {
	l.level = level
}

func (l *Logger) Writer(writer io.Writer) {
	l.writer = writer
}

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	l.logger(level, fmt.Sprintf(format, v...))
}

func (l *Logger) Log(level Level, v ...interface{}) {
	l.logger(level, fmt.Sprint(v...))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.logger(TRACE, fmt.Sprintf(format, v...))
}

func (l *Logger) Trace(v ...interface{}) {
	l.logger(TRACE, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger(DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger(DEBUG, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger(INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.logger(INFO, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger(WARN, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.logger(WARN, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.logger(ERROR, fmt.Sprint(v...))
}

func (l *Logger) logger(level Level, message string) {
	if level >= l.level {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		name := fmt.Sprintf("[%s]", level.String())

		entry := fmt.Sprintf("%s %7s :: %s\n", timestamp, name, message)
		l.writer.Write([]byte(entry))
	}
}

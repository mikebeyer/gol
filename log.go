package gol

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	level  Level
	format string
	writer []io.Writer
}

func ClassicLogger() *Logger {
	return &Logger{
		level:  INFO,
		format: time.RFC3339,
		writer: []io.Writer{os.Stderr},
	}
}

func LevelLogger(level Level) *Logger {
	return &Logger{
		level:  level,
		format: time.RFC3339,
		writer: []io.Writer{os.Stderr},
	}
}

func New(level Level, format string, writer ...io.Writer) *Logger {
	return &Logger{
		level:  level,
		format: format,
		writer: writer,
	}
}

func (l *Logger) Level(level Level) {
	l.level = level
}

func (l *Logger) Writer(writer ...io.Writer) {
	l.writer = writer
}

func (l *Logger) AddWriter(writer io.Writer) {
	l.writer = append(l.writer, writer)
}

func (l *Logger) Format(format string) {
	l.format = format
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
		timestamp := time.Now().UTC().Format(l.format)
		name := fmt.Sprintf("[%s]", level.String())

		entry := fmt.Sprintf("%s %7s :: %s\n", timestamp, name, message)

		for _, writer := range l.writer {
			writer.Write([]byte(entry))
		}
	}
}

package gol_test

import (
	"strings"
	"testing"
	"time"

	"github.com/mikebeyer/gol"
)

func TestSuppression(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.INFO, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	ammount := "12.2123123ms"
	log.Infof("%s", message)
	log.Tracef("%s", ammount)
	log.Debugf("%s", "/")

	response := writer.Data.Value[0]

	if len(writer.Data.Value) != 1 {
		t.Errorf("only 1 value should have been written, was %v", len(writer.Data.Value))
	}

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "INFO") {
		t.Errorf("expected data to contain %s was %s", "INFO", response)
	}

	if strings.Contains(response, ammount) {
		t.Errorf("expected data not to contain %s", ammount)
	}

	if strings.Contains(response, "TRACE") {
		t.Error("expected data not to contain TRACE")
	}
}

var levels = []gol.Level{
	gol.DEBUG,
	gol.TRACE,
	gol.INFO,
	gol.WARN,
	gol.ERROR,
}

func TestAll(t *testing.T) {
	for _, level := range levels {
		a := NewWriter()
		b := NewWriter()
		log := gol.New(level, time.RFC3339, a, b)

		message := "GET /apache_pb.gif HTTP/1.0 200 2326"
		log.Logf(level, "%s", message)

		respA := a.Data.Value[0]
		respB := b.Data.Value[0]

		if len(a.Data.Value) != 1 {
			t.Errorf("only 1 value should have been written, was %v", len(a.Data.Value))
		}

		if !strings.Contains(respA, message) {
			t.Errorf("expected data to contain %s was %s", message, respA)
		}

		if !strings.Contains(respA, level.String()) {
			t.Errorf("expected data to contain %s was %s", level.String(), respA)
		}

		if len(b.Data.Value) != 1 {
			t.Errorf("only 1 value should have been written, was %v", len(b.Data.Value))
		}

		if !strings.Contains(respB, message) {
			t.Errorf("expected data to contain %s was %s", message, respB)
		}

		if !strings.Contains(respB, level.String()) {
			t.Errorf("expected data to contain %s was %s", level.String(), respB)
		}
	}
}

func TestErrorf(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Errorf("%s", message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "ERROR") {
		t.Errorf("expected data to contain %s was %s", "ERROR", response)
	}
}

func TestError(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Error(message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "ERROR") {
		t.Errorf("expected data to contain %s was %s", "ERROR", response)
	}
}

func TestWarnf(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Warnf("%s", message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "WARN") {
		t.Errorf("expected data to contain %s was %s", "WARN", response)
	}
}

func TestWarn(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Warn(message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "WARN") {
		t.Errorf("expected data to contain %s was %s", "WARN", response)
	}
}

func TestInfof(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Infof("%s", message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "INFO") {
		t.Errorf("expected data to contain %s was %s", "INFO", response)
	}
}

func TestInfo(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Info(message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "INFO") {
		t.Errorf("expected data to contain %s was %s", "INFO", response)
	}
}

func TestDebugf(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Debugf("%s", message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "DEBUG") {
		t.Errorf("expected data to contain %s was %s", "DEBUG", response)
	}
}

func TestDebug(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Debug(message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "DEBUG") {
		t.Errorf("expected data to contain %s was %s", "DEBUG", response)
	}
}

func TestTracef(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Tracef("%s", message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "TRACE") {
		t.Errorf("expected data to contain %s was %s", "TRACE", response)
	}
}

func TestTrace(t *testing.T) {
	writer := NewWriter()
	log := gol.New(gol.TRACE, time.RFC3339, writer)
	message := "GET /apache_pb.gif HTTP/1.0 200 2326"
	log.Trace(message)

	response := writer.Data.Value[0]

	if !strings.Contains(response, message) {
		t.Errorf("expected data to contain %s was %s", message, response)
	}

	if !strings.Contains(response, "TRACE") {
		t.Errorf("expected data to contain %s was %s", "TRACE", response)
	}
}

type MockWriter struct {
	Data *Data
}

func NewWriter() MockWriter {
	return MockWriter{&Data{}}
}

func (m MockWriter) Write(p []byte) (n int, err error) {
	m.Data.Value = append(m.Data.Value, string(p[:]))
	return 0, nil
}

type Data struct {
	Value []string
}

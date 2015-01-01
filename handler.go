package gol

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type responseWrapper struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (r *responseWrapper) Header() http.Header {
	return r.w.Header()
}

func (r *responseWrapper) Write(b []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	size, err := r.w.Write(b)
	r.size += size
	return size, err
}

func (r *responseWrapper) WriteHeader(s int) {
	r.w.WriteHeader(s)
	r.status = s
}

func (r *responseWrapper) Status() int {
	return r.status
}

func (r *responseWrapper) Size() int {
	return r.size
}

type LoggingHandler struct {
	writer  io.Writer
	handler http.Handler
}

func NewLoggingHandler(h http.Handler, w io.Writer) LoggingHandler {
	return LoggingHandler{
		writer:  w,
		handler: h,
	}
}

func (l LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("02/Jan/2006:15:04:05 -0700")
	wrapper := &responseWrapper{
		w:      w,
		status: 200,
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	uri := r.URL.RequestURI()

	l.handler.ServeHTTP(wrapper, r)

	s := fmt.Sprintf("%s %s %s [%s] \"%s %s %s\" %v %v\n", host, "-", "-", t, r.Method, uri, r.Proto, wrapper.Status(), wrapper.Size())

	l.writer.Write([]byte(s))
}

package avitomock

import (
	"encoding/json"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	time.Sleep(time.Duration(s.conf.Delay) * time.Millisecond)
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	time.Sleep(time.Duration(s.conf.Delay) * time.Millisecond)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

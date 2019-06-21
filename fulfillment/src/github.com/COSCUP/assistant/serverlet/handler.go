package main

import (
	// "fmt"
	"github.com/COSCUP/assistant"
	"io/ioutil"
	"net/http"
)

type DebuggerWriter struct {
	w http.ResponseWriter
}

func (w DebuggerWriter) Write(d []byte) (int, error) {
	stdlog.Println("response body:", string(d))
	return w.w.Write(d)
}
func (w DebuggerWriter) Header() http.Header {
	stdlog.Println("response header:", w.Header())
	return w.w.Header()
}
func (w DebuggerWriter) WriteHeader(code int) {
	stdlog.Println("response WriteHeader:", code)
	w.w.WriteHeader(code)
}

func handler(w http.ResponseWriter, r *http.Request) {
	stdlog.Println("got connection:", r)

	if !authorizationHeaderValid(r.Header) {
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errlog.Println("request error:", err)
	}
	stdlog.Println("request body:", string(data))
	assistant.RequestHandler(DebuggerWriter{w: w}, r, data)

}

func authorizationHeaderValid(h http.Header) bool {
	return true
}

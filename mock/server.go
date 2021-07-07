package mock

import (
	"encoding/json"
	"encoding/xml"
	"net"
	"net/http"
	"test-task-20210707/types"
)

type Responses struct {
	R1 types.Response1 // response for API 1 endpoint
	R2 types.Response2 // response for API 2 endpoint
	R3 types.Response2 // response for API 3 endpoint
}

type Server struct {
	responses Responses
	listener  net.Listener
}

func New(responses Responses) *Server {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	return &Server{
		responses: responses,
		listener:  listener,
	}
}

func (s *Server) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api1", func(w http.ResponseWriter, _ *http.Request) {
		res, _ := json.Marshal(s.responses.R1)
		_, _ = w.Write(res)
	})
	mux.HandleFunc("/api2json", func(w http.ResponseWriter, _ *http.Request) {
		res, _ := json.Marshal(s.responses.R2)
		_, _ = w.Write(res)
	})
	mux.HandleFunc("/api2xml", func(w http.ResponseWriter, _ *http.Request) {
		res, _ := xml.Marshal(s.responses.R3)
		_, _ = w.Write(res)
	})
	if err := http.Serve(s.listener, mux); err != nil {
		println(err.Error())
	}
}

func (s *Server) GetPort() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}

func (s *Server) Close() error {
	return s.listener.Close()
}

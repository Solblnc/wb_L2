package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

type Port struct {
	Port string `json:"port"`
}

type Server struct {
	Cache map[string]Event
	Mu    sync.RWMutex
	Port  Port
}

func NewServer() (*Server, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	port := Port{}
	err = json.Unmarshal(data, &port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cache: make(map[string]Event),
		Port:  port,
	}, nil
}

func (s *Server) setupGetHandlers() {
	http.HandleFunc("/event_by_name", loggerMiddleware(s.EventByName))
	http.HandleFunc("/events_for_day", loggerMiddleware(s.EventsForDay))
	http.HandleFunc("/events_for_week", loggerMiddleware(s.EventsForWeek))
	http.HandleFunc("/events_for_month", loggerMiddleware(s.EventsForMonth))
}

func (s *Server) setupPostHandlers() {
	http.HandleFunc("/create_event", loggerMiddleware(s.CreateEvent))
	http.HandleFunc("/update_event", loggerMiddleware(s.UpdateEvent))
	http.HandleFunc("/delete_event", loggerMiddleware(s.DeleteEvent))
}

func (s *Server) SetupHandlers() {
	s.setupGetHandlers()
	s.setupPostHandlers()
}

func (s *Server) Run() {
	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

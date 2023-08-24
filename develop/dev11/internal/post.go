package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) setEvent(event Event) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Cache[event.Name] = event
}

//func (s *Server) deleteEvent(eventName string) {
//	s.Mu.Lock()
//	defer s.Mu.Unlock()
//	delete(s.Cache, eventName)
//}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, []byte(`error in reading a body`))
		return
	}

	if err = json.Unmarshal(data, &event); err != nil {
		jsonRespond(w, http.StatusBadRequest, []byte(`error in reading a body`))
		return
	}
	s.setEvent(event)

	log.Println(event)
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event := Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, []byte(`error in reading a body`))
		return
	}

	if err = json.Unmarshal(data, &event); err != nil {
		jsonRespond(w, http.StatusBadRequest, []byte(`error in unmarshalling`))
		return
	}

	id := event.UserID
	newId := -1

	for i, value := range s.Cache {
		if s.Cache[i].UserID == id {
			newId = event.UserID
			s.Cache[value.Name] = event
		}
	}

	if newId == -1 {
		jsonRespond(w, http.StatusBadRequest, []byte(`error in unmarshalling`))
		return
	}

	jsonRespond(w, http.StatusOK, []byte(`info is updated`))

}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event := Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		jsonRespond(w, http.StatusInternalServerError, []byte(`error in reading a body`))
		return
	}

	if err = json.Unmarshal(data, &event); err != nil {
		return
	}

	id := event.UserID

	for i, value := range s.Cache {
		if s.Cache[i].UserID == id {
			s.Cache[value.Name] = Event{}
		} else {
			jsonRespond(w, http.StatusBadRequest, []byte(`error in unmarshalling`))
		}
	}

	jsonRespond(w, http.StatusOK, []byte(`deleted`))

}

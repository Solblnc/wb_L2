package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	result := make([]Event, 0)
	s.Mu.RLock()
	for i, event := range s.Cache {
		if s.Cache[i].Date == time.Now().Format("02.01.2006") {
			result = append(result, event)
		}

	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		jsonRespond(w, http.StatusInternalServerError, []byte(`server error`))
		return
	}
	jsonRespond(w, 200, data)
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	startWeek := now.AddDate(0, 0, -int(now.Weekday()))
	endWeek := startWeek.AddDate(0, 0, 7)

	for i := range s.Cache {
		eventDate, err := time.Parse("02.01.2006", s.Cache[i].Date)
		if err != nil {
			log.Println(err)
			return
		}
		if eventDate.After(startWeek) && eventDate.Before(endWeek) {
			res, err := json.Marshal(s.Cache[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write(res)
		}
	}

}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	//now := time.Now()
	//startMonth := now.AddDate(0, 0, -int(now.Month()))
	//endMonth := startMonth.AddDate(0, 1, 0)
	//
	//for i := range s.Cache {
	//	eventDate, err := time.Parse("02.01.2006", s.Cache[i].Date)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	if eventDate.After(startMonth) && eventDate.Before(endMonth) {
	//		res, err := json.Marshal(s.Cache[i])
	//		if err != nil {
	//			log.Fatal(err)
	//			return
	//		}
	//		w.Write(res)
	//	}
	//}

	result := make([]Event, 0)
	s.Mu.RLock()
	for i, event := range s.Cache {
		if s.Cache[i].Date[3:] == time.Now().Format("02.01.2006")[3:] {
			result = append(result, event)
		}

	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		jsonRespond(w, http.StatusInternalServerError, []byte(`server error`))
		return
	}
	jsonRespond(w, 200, data)

}

func (s *Server) EventByName(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	eventName, ok := m["event_name"]
	if !ok {
		jsonRespond(w, http.StatusBadRequest, []byte(`invalid request`))
		return

	}
	s.Mu.RLock()
	event, ok := s.Cache[eventName[0]]
	s.Mu.RUnlock()
	if !ok {
		jsonRespond(w, http.StatusInternalServerError, []byte(`no event for this name`))
		return
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		jsonRespond(w, http.StatusInternalServerError, []byte(`internal server error`))
		return
	}
	jsonRespond(w, http.StatusOK, data)
}

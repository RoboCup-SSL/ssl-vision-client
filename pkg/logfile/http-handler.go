package logfile

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

func (s *Service) handle(w http.ResponseWriter, _ *http.Request, entity interface{}) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(entity)
	if err != nil {
		log.Println("Could not marshal data:", entity, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Println("Could not write response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Service) HandleGetLogFiles(w http.ResponseWriter, r *http.Request) {
	logFiles := s.GetLogFiles()
	s.handle(w, r, logFiles)
}

func (s *Service) HandleGetLogFileMetaData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := url.PathUnescape(vars["name"])
	if err != nil {
		log.Println("Could not unescape url:", vars["name"], err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	metaData, err := s.GetLogFileMetaData(name)
	if err != nil {
		log.Println("Could not get meta data:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.handle(w, r, metaData)
}

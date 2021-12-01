package logfile

import (
	"encoding/json"
	"fmt"
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
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

func readPathParam(varsKey string) (value string, err error) {
	value, err = url.PathUnescape(varsKey)
	if err != nil {
		log.Println("Could not unescape url:", varsKey, err)
		return "", err
	}
	return
}

func (s *Service) HandleGetLogFiles(w http.ResponseWriter, r *http.Request) {
	logFiles := s.GetLogFiles()
	s.handle(w, r, logFiles)
}

func (s *Service) HandleGetLogFileMetaData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := readPathParam(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

func (s *Service) HandleGetLogFileFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, err := readPathParam(vars["name"])
	if err != nil {
		badRequest(w, "Missing 'name' parameter: %s", err)
		return
	}
	messageType, err := readPathParam(vars["messageType"])
	if err != nil {
		badRequest(w, "Missing 'messageType' parameter: %s", err)
		return
	}
	requestedTime, err := readPathParam(vars["timestamp"])
	if err != nil {
		badRequest(w, "Missing 'timestamp' parameter: %s", err)
		return
	}

	logFile, err := s.getLogFile(name)
	if err != nil {
		log.Printf("Could not open log file %s: %v", name, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	messageId, err := strconv.Atoi(messageType)
	if err != nil {
		badRequest(w, "Could not parse message type '%s': %s", messageType, err)
		return
	}

	t, err := strconv.ParseInt(requestedTime, 10, 64)
	if err != nil {
		badRequest(w, "Could not parse requestedTime '%s': %s", requestedTime, err)
		return
	}
	message, err := logFile.GetFrame(persistence.MessageId(messageId), time.Unix(0, t))
	if err != nil {
		badRequest(w, "Could not get frame: %s", err)
		return
	}

	acceptHeader := r.Header.Get("Accept")
	var responseData []byte
	switch acceptHeader {
	case "application/x-protobuf":
		w.Header().Add("Content-Type", "application/x-protobuf")
		responseData = message.Message
	default:
		badRequest(w, "content type not supported: %s", acceptHeader)
		return
	}

	if _, err := w.Write(responseData); err != nil {
		log.Println("Could not write response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func badRequest(w http.ResponseWriter, format string, args ...interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	message := fmt.Sprintf(format, args)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Println("Failed to respond: ", err)
	}
}

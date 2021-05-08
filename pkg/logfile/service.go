package logfile

import (
	"github.com/pkg/errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Service struct {
	logFileFolder     string
	openLogfiles      map[string]*Logfile
	openLogfilesMutex sync.Mutex
}

func NewService(logFileFolder string) (s *Service) {
	s = new(Service)
	if strings.HasSuffix(logFileFolder, "/") {
		s.logFileFolder = logFileFolder
	} else {
		s.logFileFolder = logFileFolder + "/"
	}
	s.openLogfiles = map[string]*Logfile{}
	return
}

func (s *Service) Start() {

}

func (s *Service) Stop() {
	for name, logFile := range s.openLogfiles {
		if err := logFile.Close(); err != nil {
			log.Println("Could not close logfile:", name, err)
		}
	}
	s.openLogfiles = map[string]*Logfile{}
}

func (s *Service) getLogFile(name string) (*Logfile, error) {
	s.openLogfilesMutex.Lock()
	defer s.openLogfilesMutex.Unlock()
	if logFile, ok := s.openLogfiles[name]; ok {
		return logFile, nil
	}
	logFile, err := NewLogfile(s.logFileFolder + name)
	if err != nil {
		return nil, errors.Wrap(err, "Could not open log file: "+name)
	}
	s.openLogfiles[name] = logFile
	return logFile, nil
}

func (s *Service) GetLogFiles() []string {
	var files []string

	err := filepath.Walk(s.logFileFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".log") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Println("Could not read log file folder:", s.logFileFolder, err)
		return []string{}
	}
	var logFiles []string
	for _, file := range files {
		relativeName := strings.Replace(file, s.logFileFolder, "", 1)
		escapedName := url.PathEscape(relativeName)
		logFiles = append(logFiles, escapedName)
	}
	return logFiles
}

func (s *Service) GetLogFileMetaData(name string) (MetaData, error) {
	logFile, err := s.getLogFile(name)
	if err != nil {
		return MetaData{}, err
	}
	return logFile.GetMetaData(), nil
}

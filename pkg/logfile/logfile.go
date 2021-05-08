package logfile

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
	"log"
	"time"
)

type Logfile struct {
	reader   *persistence.Reader
	offsets  []int64
	metaData MetaData
}

func NewLogfile(filename string) (l *Logfile, err error) {
	l = new(Logfile)
	l.reader, err = persistence.NewReader(filename)
	if err != nil {
		return
	}
	if l.reader.IsIndexed() {
		l.offsets, err = l.reader.ReadIndex()
		l.metaData = l.readMetaData()
	} else {
		log.Printf("File '%v' is not indexed. Can not extract meta data.", filename)
	}
	return
}

func (l *Logfile) Close() error {
	if l.reader != nil {
		return l.reader.Close()
	}
	return nil
}

func (l *Logfile) GetMetaData() MetaData {
	return l.metaData
}

func (l *Logfile) readMetaData() (m MetaData) {
	m.NumFrames = uint32(len(l.offsets))
	first, err := l.reader.ReadMessageAt(l.offsets[0])
	if err != nil {
		log.Println("Can not read first message:", err)
		return
	}
	last, err := l.reader.ReadMessageAt(l.offsets[len(l.offsets)-1])
	if err != nil {
		log.Println("Can not read last message:", err)
		return
	}
	m.Duration = time.Duration(last.Timestamp - first.Timestamp)
	return
}

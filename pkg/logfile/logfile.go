package logfile

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
	"github.com/pkg/errors"
	"log"
	"time"
)

type Logfile struct {
	reader    *persistence.Reader
	offsetMap map[persistence.MessageId]map[time.Time]int64
	metaData  MetaData
}

func NewLogfile(filename string) (*Logfile, error) {
	l := new(Logfile)
	if reader, err := persistence.NewReader(filename); err != nil {
		return nil, err
	} else {
		l.reader = reader
	}
	if !l.reader.IsIndexed() {
		log.Printf("File '%v' is not indexed. Can not extract meta data.", filename)
		return l, nil
	}
	if offsets, err := l.reader.ReadIndex(); err != nil {
		return l, nil
	} else {
		l.offsetMap = map[persistence.MessageId]map[time.Time]int64{}
		var firstTime, lastTime *time.Time
		for _, offset := range offsets {
			if offsetTime, offsetMsgType, err := l.reader.ReadMessageTimeAndType(offset); err != nil {
				log.Println("Failed to read message time and type: ", err)
			} else {
				offsetMap, offsetMapExists := l.offsetMap[*offsetMsgType]
				if !offsetMapExists {
					offsetMap = map[time.Time]int64{}
					l.offsetMap[*offsetMsgType] = offsetMap
				}
				offsetMap[*offsetTime] = offset
				if firstTime == nil {
					firstTime = offsetTime
				}
				lastTime = offsetTime
			}
		}
		l.metaData.StartTime = *firstTime
		l.metaData.EndTime = *lastTime

		for msgType := range l.offsetMap {
			l.metaData.MessageTypes = append(l.metaData.MessageTypes, msgType)
		}
	}
	return l, nil
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

func (l *Logfile) GetFrame(messageType persistence.MessageId, t time.Time) (*persistence.Message, error) {
	if offsetMap, messageTypeExists := l.offsetMap[messageType]; !messageTypeExists {
		return nil, errors.Errorf("Message type unknown: %s", messageType.String())
	} else {
		if offset, offsetExists := offsetMap[t]; offsetExists {
			return l.reader.ReadMessageAt(offset)
		}
		offset := findNearestOffset(offsetMap, t)
		return l.reader.ReadMessageAt(offset)
	}
}

func findNearestOffset(offsetMap map[time.Time]int64, t time.Time) int64 {
	var prevTime time.Time
	var prevOffset int64
	for offsetTime, offset := range offsetMap {
		if offsetTime.After(t) {
			if offsetTime.Sub(t) > t.Sub(prevTime) {
				return prevOffset
			}
			return offset
		}
		prevTime = t
		prevOffset = offset
	}
	return prevOffset
}

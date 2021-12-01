package logfile

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
	"time"
)

type MetaData struct {
	StartTime    time.Time               `json:"start-time"`
	EndTime      time.Time               `json:"end-time"`
	MessageTypes []persistence.MessageId `json:"message-ids"`
}

type MetaDataList struct {
	MetaData []MetaData `json:"meta_data"`
}

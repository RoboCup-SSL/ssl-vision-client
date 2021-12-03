package logfile

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
)

type MetaData struct {
	StartTime    int64                   `json:"start-time"`
	EndTime      int64                   `json:"end-time"`
	MessageTypes []persistence.MessageId `json:"message-ids"`
}

type MetaDataList struct {
	MetaData []MetaData `json:"meta_data"`
}

package logfile

import "time"

type MetaData struct {
	NumFrames uint32        `json:"num_frames"`
	Duration  time.Duration `json:"duration"`
}

type MetaDataList struct {
	MetaData []MetaData `json:"meta_data"`
}

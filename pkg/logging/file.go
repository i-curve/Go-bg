package logging

import (
	"fmt"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = ""
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		LogSaveName,
		time.Now().Format(TimeFormat),
		LogFileExt,
	)
}

package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"template/pkg/file"
	"template/pkg/setting"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"INFO", "DEBUG", "WARN", "ERROR", "FATAL"}
)

const (
	INFO Level = iota
	DEBUG
	WARNING
	ERROR
	FATAL
)

func Setup() {
	if setting.AppSetting.LogSavePath != "" {
		LogSavePath = setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
	}
	if setting.AppSetting.LogSaveName != "" {
		LogSaveName = setting.AppSetting.LogSaveName
	}
	if setting.AppSetting.LogFileExt != "" {
		LogFileExt = setting.AppSetting.LogFileExt
	}
	if setting.ServerSetting.RunMode == "release" {
		TimeFormat = setting.AppSetting.TimeFormat
	}

	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v...)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog(v bool) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
	if v {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func getFields(args ...interface{}) log.Fields {
	fields := make(log.Fields)
	size := len(args)

	for index := 0; index < size-1; index += 2 {
		fields[args[index].(string)] = args[index+1]
	}
	return fields
}

func Debug(info string, args ...interface{}) {
	fields := getFields(args...)
	log.WithFields(fields).Debug(info)
}

func Info(info string, args ...interface{}) {
	fields := getFields(args...)
	log.WithFields(fields).Info(info)
}

func Warn(info string, args ...interface{}) {
	fields := getFields(args...)
	log.WithFields(fields).Warn(info)
}

func Error(info string, args ...interface{}) {
	fields := getFields(args...)
	log.WithFields(fields).Error(info)
}

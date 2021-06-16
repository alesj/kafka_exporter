package main

import (
	"fmt"
	stdLog "log"
)

var errorLog, infoLog *stdLog.Logger

func init() {
	stdLog.SetPrefix("[kafka_exporter] ")
	errorLog = stdLog.New(stdLog.Writer(), stdLog.Prefix()+"[ERROR] ", stdLog.Flags())
	infoLog = stdLog.New(stdLog.Writer(), stdLog.Prefix()+"[INFO] ", stdLog.Flags())
	stdLog.SetPrefix(stdLog.Prefix() + "[DEBUG] ")
}

func Errorf(s string, v ...interface{}) {
	_ = errorLog.Output(2, fmt.Sprintf(s, v...))
}

func Infof(s string, v ...interface{}) {
	_ = infoLog.Output(2, fmt.Sprintf(s, v...))
}

func Debugf(s string, v ...interface{}) {
	_ = stdLog.Output(2, fmt.Sprintf(s, v...))
}

package logger

import (
	"fmt"
	"log"
	"os"
)

//LogFile global logfile variable of type *os.File
var LogFile *os.File

//Logging global scope struct
type Logging struct {
	LogEvent *log.Logger
	LogErr   *log.Logger
	LogAlarm *log.Logger
}

//Event method set for Logging struct to log events
func (logger *Logging) Event(msg string) {
	logger.LogEvent.Println(msg)

}

//Err method set for Logging struct to log errors
func (logger *Logging) Err(msg error) {
	logger.LogErr.Panic(msg)

}

//Alarm methos set for Logging struct to log alarms
func (logger *Logging) Alarm(msg string) {
	logger.LogAlarm.Println(msg)
}

//LoggerInit initializes the log file for data logging event, alarms and errors
func LoggerInit() (logStruct Logging) {
	LogFile, err := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logStruct = Logging{
		log.New(LogFile, "Event ", log.LstdFlags), // event logger
		log.New(LogFile, "Error ", log.LstdFlags), // errors logger
		log.New(LogFile, "Alarm ", log.LstdFlags), // alarms logger
	}
	return logStruct
}

//LoggerClose closes the logger file
func LoggerClose() {
	LogFile.Close()
}

/*
FileLogger, DatabaseLogger and AdapterLogger are implementations of ILogger.
You can think ThirdPartyLogger as an imported library/package.
This pattern allows you to easily adapt 3rd party code to your code
without any breaking changes.
*/

package main

import (
	"fmt"
)

func main() {
	var logger ILogger
	logger = &FileLogger{}
	logger.writeLog("test message")
	logger = &DatabaseLogger{}
	logger.writeLog("test message")
	logger = &AdapterLogger{}
	logger.writeLog("The 3rd party logger adapted to the project by AdapterLogger.")
}

type ILogger interface {
	writeLog(message string)
}

type FileLogger struct{}

func (logger *FileLogger) writeLog(message string) {
	fmt.Println("Message: ", message, " logged to file.")
}

type DatabaseLogger struct{}

func (logger *DatabaseLogger) writeLog(message string) {
	fmt.Println("Message: ", message, " logged to database.")
}

type AdapterLogger struct {
	innerLogger ThirdPartyLogger
}

func (logger *AdapterLogger) writeLog(message string) {
	logger.innerLogger.logMessage = message
	logger.innerLogger.logTheMessage()
}

type ThirdPartyLogger struct {
	logMessage string
}

func (logger *ThirdPartyLogger) logTheMessage() {
	fmt.Println("Log: ", logger.logMessage)
}

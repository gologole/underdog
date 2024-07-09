// logger/logger.go

package logger

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
)

func Init(debugFile, infoFile string) {
	debugLogFile, err := os.OpenFile(debugFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open debug log file: %v", err)
	}
	// log.Llongfile добавляет полный путь к файлу и номер строки
	Debug = log.New(debugLogFile, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)

	infoLogFile, err := os.OpenFile(infoFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open info log file: %v", err)
	}
	Info = log.New(io.MultiWriter(infoLogFile, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

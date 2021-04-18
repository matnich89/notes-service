package logger

import (
	"log"
	"os"
)

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		errorLog: log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime),
		infoLog:  log.New(os.Stderr, "INFO ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) OutputInfo(text string) {
	l.infoLog.Println(text)
}

func (l *Logger) OutputError(text string) {
	l.errorLog.Println(text)
}

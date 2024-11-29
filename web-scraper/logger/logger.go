package main

import (
	"fmt"
	"os"
)

type LogMethods interface {
	Log(logLevel uint, message string)
	SetMinLogLevel(logLevel uint)
}

const (
	Debug = iota
	Info
	Warn
	Error
	Disabled = 100
)

type FileLogger struct {
	minLogLevel uint
	file        *os.File
}

func (log *FileLogger) Log(logLevel uint, message string) {
	if log.minLogLevel <= logLevel {
		fmt.Fprintln(log.file, message)
	}
}

func (log *FileLogger) SetMinLogLevel(logLevel uint) {
	log.minLogLevel = logLevel
}

type ConsoleLogger struct {
	minLogLevel uint
}

func (log *ConsoleLogger) Log(logLevel uint, message string) {
	if log.minLogLevel <= logLevel {
		fmt.Println(message)
	}
}

func (log *ConsoleLogger) SetMinLogLevel(logLevel uint) {
	log.minLogLevel = logLevel
}

type MultiLogger struct {
	loggers []LogMethods
}

func (log *MultiLogger) Log(logLevel uint, message string) {
	for _, logger := range log.loggers {
		logger.Log(logLevel, message)
	}
}

func main() {
	f1, err := os.OpenFile("./web-scraper/logger/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log := &MultiLogger{
		loggers: []LogMethods{
			&ConsoleLogger{
				minLogLevel: Info,
			},
			&FileLogger{
				minLogLevel: Warn,
				file:        f1,
			},
		},
	}

	log.Log(Info, "Hello")
	log.Log(Warn, "Oops")
	log.Log(Error, "Errored")
}

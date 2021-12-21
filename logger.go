package logger

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	fatalLog = log.New(os.Stdout, "\033[31;3m[FATAL]\033[0m ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\033[33m[WARN]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "\033[34m[DEBUG]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{fatalLog, errorLog, warnLog, infoLog, debugLog}

	mu sync.Mutex
)

var (
	Fatal  = fatalPrintln
	Fatalf = fatalPrintf
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Warn   = warnLog.Println
	Warnf  = warnLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if FatalLevel < level {
		fatalLog.SetOutput(ioutil.Discard)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if WarnLevel < level {
		warnLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}

	if DebugLevel < level {
		debugLog.SetOutput(ioutil.Discard)
	}
}

func fatalPrintln(v ...interface{}) {
	fatalLog.Println(v...)
	os.Exit(1)
}

func fatalPrintf(format string, v ...interface{}) {
	fatalLog.Printf(format, v...)
	os.Exit(1)
}

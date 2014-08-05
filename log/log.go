// Package log adds helper functions to standard logger
package log

import (
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type Logger struct {
	*log.Logger
}

// New creates a new Logger.
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{log.New(out, prefix, flag)}
}

// CheckFatal writes a log and exit the program if err is not nil.
func (l *Logger) CheckFatal(err error, v ...interface{}) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		l.Fatalf("%s:%d: "+format(v)+" - %s\n", append(append([]interface{}{filepath.Base(file), line}, v...), err)...)
	}
}

// CheckLog writes a log if err is not nil.
func (l *Logger) CheckLog(err error, v ...interface{}) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		l.Printf("%s:%d: "+format(v)+" - %s\n", append(append([]interface{}{filepath.Base(file), line}, v...), err)...)
	}
}

// Log writes a log.
func (l *Logger) Log(v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Printf("%s:%d: "+format(v)+"\n", append([]interface{}{filepath.Base(file), line}, v...)...)
}

func format(v []interface{}) string {
	if len(v) == 0 {
		return ""
	}

	sym := "%s"
	return sym + strings.Repeat(" "+sym, len(v)-1)
}

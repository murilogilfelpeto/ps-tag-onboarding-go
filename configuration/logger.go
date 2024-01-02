package configuration

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", logger.Flags()),
		info:   log.New(writer, "INFO: ", logger.Flags()),
		warn:   log.New(writer, "WARN: ", logger.Flags()),
		error:  log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Debug(value ...interface{}) {
	l.debug.Println(value...)
}

func (l *Logger) Info(value ...interface{}) {
	l.info.Println(value...)
}

func (l *Logger) Warn(value ...interface{}) {
	l.warn.Println(value...)
}

func (l *Logger) Error(value ...interface{}) {
	l.error.Println(value...)
}

func (l *Logger) Debugf(format string, value ...interface{}) {
	l.debug.Printf(format, value...)
}

func (l *Logger) Infof(format string, value ...interface{}) {
	l.info.Printf(format, value...)
}

func (l *Logger) Warnf(format string, value ...interface{}) {
	l.warn.Printf(format, value...)
}

func (l *Logger) Errorf(format string, value ...interface{}) {
	l.error.Printf(format, value...)
}

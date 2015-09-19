package log

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	// meta information for a logger.
	meta = logrus.Fields{}

	// appLogger for debug, info, warn.
	appLogger = &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{},
	}

	// errLogger for error, panic.
	errLogger = &logrus.Logger{
		Out:       os.Stderr,
		Level:     logrus.ErrorLevel,
		Formatter: &logrus.JSONFormatter{},
	}
)

// AddMetaInfo adds meta information.
func AddMetaInfo(f logrus.Fields) {
	for k, v := range f {
		meta[k] = v
	}
}

// SetAppLogger sets a logger for debug, info, warn.
func SetAppLogger(l *logrus.Logger) {
	appLogger = l
}

// SetErrLogger sets a logger for error, panic.
func SetErrLogger(l *logrus.Logger) {
	errLogger = l
}

// Debug messages writes on Stdout by async.
func Debug(msg string, fs ...logrus.Fields) {
	go joinFields(appLogger.WithFields(meta), fs).Debugln(msg)
}

// Info messages writes on Stdout by async.
func Info(msg string, fs ...logrus.Fields) {
	go joinFields(appLogger.WithFields(meta), fs).Infoln(msg)
}

// Warn messages writes on Stdout by async.
func Warn(msg string, fs ...logrus.Fields) {
	go joinFields(appLogger.WithFields(meta), fs).Warnln(msg)
}

// Error messages writes on Stderr by async.
func Error(msg string, err error, fs ...logrus.Fields) {
	go joinFields(errLogger.WithFields(meta), fs).WithError(err).Errorln(msg)
}

// Panic messages writes on Stderr.
func Panic(msg string, err error, fs ...logrus.Fields) {
	// ** DON'T ASYNC **
	joinFields(errLogger.WithFields(meta), fs).WithError(err).Panicln(msg)
}

// joinFields on a entry.
func joinFields(e *logrus.Entry, fs []logrus.Fields) *logrus.Entry {

	if len(fs) == 0 {
		return e
	}

	for i := range fs {
		e = e.WithFields(fs[i])
	}

	return e
}

package log

import "github.com/Sirupsen/logrus"

var (
	// metadata for a logger.
	metadata = logrus.Fields{}

	// logger for application.
	logger = logrus.New()
)

// AddMetadata adds metadata.
func AddMetadata(f logrus.Fields) {
	for k, v := range f {
		metadata[k] = v
	}
}

// SetLogger sets a logger for application.
func SetLogger(l *logrus.Logger) {
	logger = l
}

// AddHooks to a logger.
func AddHooks(h ...logrus.Hook) {

	if len(h) == 0 {
		return
	}

	for i := range h {
		logger.Hooks.Add(h[i])
	}
}

// Debug messages writes on logger.Out by async.
func Debug(msg string, fs ...logrus.Fields) {
	go joinFields(logger.WithFields(metadata), fs).Debugln(msg)
}

// Info messages writes on logger.Out by async.
func Info(msg string, fs ...logrus.Fields) {
	go joinFields(logger.WithFields(metadata), fs).Infoln(msg)
}

// Warn messages writes on logger.Out by async.
func Warn(msg string, fs ...logrus.Fields) {
	go joinFields(logger.WithFields(metadata), fs).Warnln(msg)
}

// Error messages writes on logger.Out by async.
func Error(msg string, err error, fs ...logrus.Fields) {
	go joinFields(logger.WithFields(metadata), fs).WithError(err).Errorln(msg)
}

// Fatal messages writes on logger.Out.
func Fatal(msg string, err error, fs ...logrus.Fields) {
	// ** DON'T ASYNC **
	joinFields(logger.WithFields(metadata), fs).WithError(err).Fatalln(msg)
}

// Panic messages writes on logger.Out.
func Panic(msg string, err error, fs ...logrus.Fields) {
	// ** DON'T ASYNC **
	joinFields(logger.WithFields(metadata), fs).WithError(err).Panicln(msg)
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

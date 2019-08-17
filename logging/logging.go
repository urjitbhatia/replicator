package logging

import (
	log "github.com/Sirupsen/logrus"
)

// SetLevel sets the log level.
// Accepted levels are panic, fatal, error, warn, info and debug.
func SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Fatalf(`not a valid level: "%s"`, level)
	}
	log.SetLevel(lvl)
}

// Debug logs a message with severity DEBUG.
func Debug(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Info logs a message with severity INFO.
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warning logs a message with severity WARNING.
func Warning(format string, v ...interface{}) {
	log.Warningf(format, v...)
}

// Error logs a message with severity ERROR.
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

// Fatal logs a message with severity ERROR which is then followed by a call
// to os.Exit().
func Fatal(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

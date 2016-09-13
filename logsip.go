// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"errors"
	"fmt"
	"time"
)

type Level uint8

var startTime time.Time

func init() {
	startTime = time.Now()
}

func sinceStartTime() int {
	return int(time.Since(startTime) / time.Second)
}

func (level Level) String() string {
	switch level {
	case PanicLevel:
		return fmt.Sprintf(Colorize("{{.Red}}PANIC[%04d] ==> {{.Default}}"), sinceStartTime())
	case FatalLevel:
		return fmt.Sprintf(Colorize("{{.Red}}FATAL[%04d] ==> {{.Default}}"), sinceStartTime())
	case ErrorLevel:
		return fmt.Sprintf(Colorize("{{.Red}}ERROR[%04d] ==> {{.Default}}"), sinceStartTime())
	case WarnLevel:
		return fmt.Sprintf(Colorize("{{.Yellow}}WARN[%04d] ==> {{.Default}}"), sinceStartTime())
	case InfoLevel:
		return fmt.Sprintf(Colorize("{{.Blue}}INFO[%04d] ==> {{.Default}}"), sinceStartTime())
	case DebugLevel:
		return fmt.Sprintf(Colorize("{{.Purple}}DEBUG[%04d] ==> {{.Default}}"), sinceStartTime())
	}

	return "[UNKNOWN] ==> "
}

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func parseLevel(level Level) (Level, error) {
	switch level {
	case PanicLevel:
		return PanicLevel, nil
	case FatalLevel:
		return FatalLevel, nil
	case ErrorLevel:
		return ErrorLevel, nil
	case WarnLevel:
		return WarnLevel, nil
	case InfoLevel:
		return InfoLevel, nil
	case DebugLevel:
		return DebugLevel, nil
	}

	return PanicLevel, errors.New("unable to parse log level")
}

func (logger *Logger) SetLevel(level Level) {
	l, err := parseLevel(level)
	if err != nil {
		fmt.Printf("%v", err)
	}

	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Level = l
}

func (logger *Logger) GetLevel() Level {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.Level
}

type Logging interface {
	Panic(v ...interface{})
	Fatal(v ...interface{})
	Error(v ...interface{})
	Warn(v ...interface{})
	Info(v ...interface{})
	Debug(v ...interface{})

	Panicf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})

	Panicln(v ...interface{})
	Fatalln(v ...interface{})
	Errorln(v ...interface{})
	Warnln(v ...interface{})
	Infoln(v ...interface{})
	Debugln(v ...interface{})
}

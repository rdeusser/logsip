// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/iamthemuffinman/logsip/isatty"
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
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Red}}PANIC[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("PANIC[%04d] ", sinceStartTime())
	case FatalLevel:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Red}}FATAL[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("FATAL[%04d] ", sinceStartTime())
	case ErrorLevel:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Red}}ERROR[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("ERROR[%04d] ", sinceStartTime())
	case WarnLevel:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Yellow}}WARN[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("WARN[%04d] ", sinceStartTime())
	case InfoLevel:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Blue}}INFO[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("INFO[%04d] ", sinceStartTime())
	case DebugLevel:
		if isatty.IsTerminal(os.Stdout.Fd()) {
			return fmt.Sprintf(prettyPrint("{{.Purple}}DEBUG[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("DEBUG[%04d] ", sinceStartTime())
	}

	return fmt.Sprintf("UNKNOWN ")
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

func SetLevel(level Level) {
	l, err := parseLevel(level)
	if err != nil {
		fmt.Printf("%v", err)
	}

	std.mu.Lock()
	defer std.mu.Unlock()
	std.Level = l
}

func GetLevel() Level {
	std.mu.Lock()
	defer std.mu.Unlock()
	return std.Level
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

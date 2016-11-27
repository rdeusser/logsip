// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/iamthemuffinman/logsip/isatty"
)

// Level is a type that represents the log level
type Level uint8

var startTime time.Time

// IsTerminal defines whether we have a terminal or not
var IsTerminal = isatty.IsTerminal(os.Stdout.Fd())

const (
	// PanicLevel is the lowest level
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func init() {
	startTime = time.Now()
}

func sinceStartTime() int {
	return int(time.Since(startTime) / time.Second)
}

func (level Level) String() string {
	switch level {
	case PanicLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Red}}PANIC[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("PANIC[%04d] ", sinceStartTime())
	case FatalLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Red}}FATAL[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("FATAL[%04d] ", sinceStartTime())
	case ErrorLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Red}}ERROR[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("ERROR[%04d] ", sinceStartTime())
	case WarnLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Yellow}}WARN[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("WARN[%04d] ", sinceStartTime())
	case InfoLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Blue}}INFO[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("INFO[%04d] ", sinceStartTime())
	case DebugLevel:
		if IsTerminal {
			return fmt.Sprintf(prettyPrint("{{.Purple}}DEBUG[%04d]{{.Default}} "), sinceStartTime())
		}
		return fmt.Sprintf("DEBUG[%04d] ", sinceStartTime())
	}

	return fmt.Sprintf("UNKNOWN ")
}

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

// SetLevel is a method to set the log level (i.e Info, Warn, Debug, etc.)
func (logger *Logger) SetLevel(level Level) {
	l, err := parseLevel(level)
	if err != nil {
		fmt.Printf("%v", err)
	}

	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Level = l
}

// GetLevel is a method to get the current log level
func (logger *Logger) GetLevel() Level {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.Level
}

// SetLevel is a function to set the log level (i.e Info, Warn, Debug, etc.)
func SetLevel(level Level) {
	l, err := parseLevel(level)
	if err != nil {
		fmt.Printf("%v", err)
	}

	std.mu.Lock()
	defer std.mu.Unlock()
	std.Level = l
}

// GetLevel is a function to get the current log level
func GetLevel() Level {
	std.mu.Lock()
	defer std.mu.Unlock()
	return std.Level
}

// SetOutput allows you to use the global logger with any io.Writer
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

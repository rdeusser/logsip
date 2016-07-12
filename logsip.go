// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger is a modified log.Logger which provides logging levels.
type Logger struct {
	WarnPrefix  string
	FatalPrefix string
	InfoPrefix  string
	PanicPrefix string
	DebugPrefix string
	DebugMode   bool
	*log.Logger
}

type Log interface {
	Info(v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})

	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})

	Infoln(format string, v ...interface{})
	Debugln(format string, v ...interface{})
	Warnln(format string, v ...interface{})
	Fatalln(format string, v ...interface{})
	Panicln(format string, v ...interface{})
}

// New returns the Default logger, but you can specify anything that satisifes the io.Writer interface.
func New(out io.Writer, pkg string) *Logger {
	return &Logger{
		WarnPrefix:  Colorize("{{.Yellow}}==> WARN:{{.Default}} "),
		InfoPrefix:  Colorize("{{.Green}}==> INFO:{{.Default}} "),
		FatalPrefix: Colorize("{{.Red}}==> FATAL:{{.Default}} "),
		PanicPrefix: Colorize("{{.Red}}==> PANIC:{{.Default}} "),
		DebugPrefix: Colorize("{{.Cyan}}==> DEBUG:{{.Default}} "),
		DebugMode:   false,
		Logger:      log.New(out, "["+pkg+"] : ", 0),
	}
}

// Default returns a new Logger with the default color and prefix options.
func Default() *Logger {
	return &Logger{
		WarnPrefix:  Colorize("{{.Yellow}}==> WARN:{{.Default}} "),
		InfoPrefix:  Colorize("{{.Green}}==> INFO:{{.Default}} "),
		FatalPrefix: Colorize("{{.Red}}==> FATAL:{{.Default}} "),
		PanicPrefix: Colorize("{{.Red}}==> PANIC:{{.Default}} "),
		DebugPrefix: Colorize("{{.Cyan}}==> DEBUG:{{.Default}} "),
		DebugMode:   false,
		Logger:      log.New(os.Stdout, "", 0),
	}
}

// Info works just like log.Print, but with a Green "==> INFO:" prefix.
func (l *Logger) Info(v ...interface{}) {
	l.Logger.Print(l.InfoPrefix + fmt.Sprint(v...))
}

// Debug works just like log.Print, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debug(v ...interface{}) {
	if l.DebugMode {
		l.Logger.Print(l.DebugPrefix + fmt.Sprint(v...))
	}
}

// Warn works just like log.Print, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warn(v ...interface{}) {
	l.Logger.Print(l.WarnPrefix + fmt.Sprint(v...))
}

// Fatal works just like log.Fatal, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(l.FatalPrefix + fmt.Sprint(v...))
}

// Panic works just like log.Panic, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panic(v ...interface{}) {
	l.Logger.Panic(l.PanicPrefix + fmt.Sprint(v...))
}

// Infof works just like log.Printf, but with a Green "==> INFO:" prefix.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logger.Print(l.InfoPrefix + fmt.Sprintf(format, v...))
}

// Debugf works just like log.Printf, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.DebugMode {
		l.Logger.Print(l.DebugPrefix + fmt.Sprintf(format, v...))
	}
}

// Warnf works just like log.Printf, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logger.Print(l.WarnPrefix + fmt.Sprintf(format, v...))
}

// Fatalf works just like log.Fatalf, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatalf(l.FatalPrefix + fmt.Sprintf(format, v...))
}

// Panicf works just like log.Panicf, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(l.PanicPrefix + fmt.Sprintf(format, v...))
}

// Infoln works just like log.Println, but with a Green "==> INFO:" prefix.
func (l *Logger) Infoln(v ...interface{}) {
	l.Logger.Println(l.InfoPrefix + fmt.Sprint(v...))
}

// Debugln works just like log.Println, but with a Cyan "==> DEBUG:" prefix.
func (l *Logger) Debugln(v ...interface{}) {
	if l.DebugMode {
		l.Logger.Println(l.DebugPrefix + fmt.Sprint(v...))
	}
}

// Warnln works just like log.Println, but with a Yellow "==> WARN:" prefix.
func (l *Logger) Warnln(v ...interface{}) {
	l.Logger.Println(l.WarnPrefix + fmt.Sprint(v...))
}

// Fatalln works just like log.Fatalln, but with a Red "==> FATAL:" prefix.
func (l *Logger) Fatalln(v ...interface{}) {
	l.Logger.Fatalln(l.FatalPrefix + fmt.Sprint(v...))
}

// Panicln works just like log.Panicln, but with a Red "==> PANIC:" prefix.
func (l *Logger) Panicln(v ...interface{}) {
	l.Logger.Panicln(l.PanicPrefix + fmt.Sprint(v...))
}

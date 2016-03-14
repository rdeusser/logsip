// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

// Logger is a modified log.Logger which provides logging levels.
type Logger struct {
	WarnPrefix  string
	FatalPrefix string
	InfoPrefix  string
	PanicPrefix string
	*log.Logger
}

// New returns a blank Logger so you can define your own prefixes.
func New(out io.Writer) *Logger {
	return &Logger{Logger: log.New(out, "", 0)}
}

// Default returns a new Logger with the default color and prefix options.
func Default() *Logger {
	return &Logger{
		WarnPrefix:  color.New(color.FgYellow).SprintFunc()("==> Warn: "),
		InfoPrefix:  color.New(color.FgCyan).SprintFunc()("==> Info: "),
		FatalPrefix: color.New(color.FgRed).SprintFunc()("==> Fatal: "),
		PanicPrefix: color.New(color.FgRed).SprintFunc()("==> Panic: "),
		Logger:      log.New(os.Stdout, "", 0),
	}
}

// Info works just like log.Print however adds the Info prefix.
func (l *Logger) Info(v ...interface{}) {
	l.Print(l.InfoPrefix + fmt.Sprint(v...))
}

// Warn works just like log.Print however adds the Warn prefix.
func (l *Logger) Warn(v ...interface{}) {
	l.Print(l.WarnPrefix + fmt.Sprint(v...))
}

// Fatal works just like log.Fatal however adds the Fatal prefix.
func (l *Logger) Fatal(v ...interface{}) {
	l.Fatal(l.FatalPrefix + fmt.Sprint(v...))
}

// Panic works just like log.Panic however adds the Panic prefix.
func (l *Logger) Panic(v ...interface{}) {
	l.Panic(l.PanicPrefix + fmt.Sprint(v...))
}

// Infof works just like log.Printf however adds the Info prefix.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Print(l.InfoPrefix + fmt.Sprintf(format, v...))
}

// Warnf works just like log.Printf however adds the Warn prefix.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Print(l.WarnPrefix + fmt.Sprintf(format, v...))
}

// Fatalf works just like log.Fatalf however adds the Fatal prefix.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Fatalf(l.FatalPrefix + fmt.Sprintf(format, v...))
}

// Panicf works just like log.Panicf however adds the Panic prefix.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Panicf(l.PanicPrefix + fmt.Sprintf(format, v...))
}

// Infoln works just like log.Println however adds the Info prefix
func (l *Logger) Infoln(v ...interface{}) {
	l.Println(l.InfoPrefix + fmt.Sprint(v...))
}

// Warnln works just like log.Println however adds the Warnln prefix
func (l *Logger) Warnln(v ...interface{}) {
	l.Println(l.WarnPrefix + fmt.Sprint(v...))
}

// Fatalln works just like log.Fatalln however adds the Fatal prefix
func (l *Logger) Fatalln(v ...interface{}) {
	l.Fatalln(l.FatalPrefix + fmt.Sprint(v...))
}

// Panicln works just like log.Panicln however adds the Panic prefix
func (l *Logger) Panicln(v ...interface{}) {
	l.Panicln(l.PanicPrefix + fmt.Sprint(v...))
}

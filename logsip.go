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

type Logger struct {
	WarnPrefix  string
	FatalPrefix string
	InfoPrefix  string
	PanicPrefix string
	*log.Logger
}

func New(out io.Writer) *Logger {
	return &Logger{Logger: log.New(out, "", 0)}
}

func Default() *Logger {
	return &Logger{
		WarnPrefix:  color.New(color.FgYellow).SprintFunc()("==> Warn: "),
		InfoPrefix:  color.New(color.FgCyan).SprintFunc()("==> Info: "),
		FatalPrefix: color.New(color.FgRed).SprintFunc()("==> Fatal: "),
		PanicPrefix: color.New(color.FgRed).SprintFunc()("==> Panic: "),
		Logger:      log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.Print(l.InfoPrefix + fmt.Sprint(v...))
}
func (l *Logger) Warn(v ...interface{}) {
	l.Print(l.WarnPrefix + fmt.Sprint(v...))
}
func (l *Logger) Fatal(v ...interface{}) {
	l.Print(l.FatalPrefix + fmt.Sprint(v...))
	os.Exit(1)
}
func (l *Logger) Panic(v ...interface{}) {
	l.Print(l.PanicPrefix + fmt.Sprint(v...))
	panic(l)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Print(l.InfoPrefix + fmt.Sprintf(format, v...))
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Print(l.WarnPrefix + fmt.Sprintf(format, v...))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Print(l.FatalPrefix + fmt.Sprintf(format, v...))
	os.Exit(1)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Print(l.PanicPrefix + fmt.Sprintf(format, v...))
	panic(l)
}
func (l *Logger) Infoln(v ...interface{}) {
	l.Println(l.InfoPrefix + fmt.Sprint(v...))
}
func (l *Logger) Warnln(v ...interface{}) {
	l.Println(l.WarnPrefix + fmt.Sprint(v...))
}
func (l *Logger) Fatalln(v ...interface{}) {
	l.Println(l.FatalPrefix + fmt.Sprint(v...))
	os.Exit(1)
}
func (l *Logger) Panicln(v ...interface{}) {
	l.Println(l.PanicPrefix + fmt.Sprint(v...))
	panic(l)
}

// All this package really does is add a little structure and a little color
// to the standard log package

package logsip

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	Out io.Writer
}

func New(Out io.Writer) *Logger {
	return &Logger{Out: Out}
}

var cyan = color.New(color.FgCyan).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func (l *Logger) Info(v ...interface{}) {
	prefix := cyan("==> Info: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Print(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	prefix := yellow("==> Warn: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Print(v...)
}
func (l *Logger) Fatal(v ...interface{}) {
	prefix := red("==> Fatal: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Print(v...)
	os.Exit(1)
}
func (l *Logger) Panic(v ...interface{}) {
	prefix := red("==> Panic: ")
	logger := log.New(l.Out, prefix+"-- ", log.Llongfile)
	logger.Print(v...)
	panic(logger)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	prefix := cyan("==> Info: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	prefix := yellow("==> Warn: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Printf(format, v...)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	prefix := red("==> Fatal: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Printf(format, v...)
	os.Exit(1)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	prefix := red("==> Panic: ")
	logger := log.New(l.Out, prefix+"-- ", log.Llongfile)
	logger.Printf(format, v...)
	panic(logger)
}

func (l *Logger) Infoln(v ...interface{}) {
	prefix := cyan("==> Info: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Println(v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	prefix := yellow("==> Warn: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Println(v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	prefix := red("==> Fatal: ")
	logger := log.New(l.Out, prefix, 0)
	logger.Println(v...)
	os.Exit(1)
}

func (l *Logger) Panicln(v ...interface{}) {
	prefix := red("==> Panic: ")
	logger := log.New(l.Out, prefix+"-- ", log.Llongfile)
	logger.Println(v...)
	panic(logger)
}

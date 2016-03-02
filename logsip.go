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
	out io.Writer
	*log.Logger
}

func New(out io.Writer) *Logger {
	return &Logger{out: out}
}

var cyan = color.New(color.FgCyan).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func (l *Logger) Info(v ...interface{}) {
	prefix := cyan("==> Info: ")
	l.Print(prefix + fmt.Sprint(v...))
}
func (l *Logger) Warn(v ...interface{}) {
	prefix := yellow("==> Warn: ")
	l.Print(prefix + fmt.Sprint(v...))
}
func (l *Logger) Fatal(v ...interface{}) {
	prefix := red("==> Fatal: ")
	l.Print(prefix + fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	prefix := red("==> Panic: ")
	l.Print(prefix + fmt.Sprint(v...))
	panic(l)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	prefix := cyan("==> Info: ") + format
	l.Printf(prefix + fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	prefix := yellow("==> Warn: ") + format
	l.Printf(prefix + fmt.Sprint(v...))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	prefix := red("==> Fatal: ") + format
	l.Printf(prefix + fmt.Sprint(v...))
	os.Exit(1)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	prefix := red("==> Panic: ") + format
	l.Printf(prefix + fmt.Sprint(v...))
	panic(l)
}

func (l *Logger) Infoln(v ...interface{}) {
	prefix := cyan("==> Info: ")
	l.Println(prefix + fmt.Sprint(v...))
}

func (l *Logger) Warnln(v ...interface{}) {
	prefix := yellow("==> Warn: ")
	l.Println(prefix + fmt.Sprint(v...))
}

func (l *Logger) Fatalln(v ...interface{}) {
	prefix := red("==> Fatal: ")
	l.Println(prefix + fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Panicln(v ...interface{}) {
	prefix := red("==> Panic: ")
	l.Println(prefix + fmt.Sprint(v...))
	panic(l)
}

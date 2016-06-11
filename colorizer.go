// Inspiration (and some code) taken from https://github.com/deis/pkg/prettyprint
// The idea is not to rely on anything except the standard library

package logsip

import (
	"bytes"
	"html/template"
)

var Colors = map[string]string{
	"Default":     "\033[0m",
	"Black":       "\033[0;30m",
	"Red":         "\033[0;31m",
	"Green":       "\033[0;32m",
	"Yellow":      "\033[0;33m",
	"Blue":        "\033[0;34m",
	"Purple":      "\033[0;35m",
	"Cyan":        "\033[0;36m",
	"White":       "\033[0;37m",
	"BoldBlack":   "\033[1;30m",
	"BoldRed":     "\033[1;31m",
	"BoldGreen":   "\033[1;32m",
	"BoldYellow":  "\033[1;33m",
	"BoldBlue":    "\033[1;34m",
	"BoldPurple":  "\033[1;35m",
	"BoldCyan":    "\033[1;36m",
	"BoldWhite":   "\033[1;37m",
	"UnderBlack":  "\033[4;30m",
	"UnderRed":    "\033[4;31m",
	"UnderGreen":  "\033[4;32m",
	"UnderYellow": "\033[4;33m",
	"UnderBlue":   "\033[4;34m",
	"UnderPurple": "\033[4;35m",
	"UnderCyan":   "\033[4;36m",
	"UnderWhite":  "\033[4;37m",
	"HiBlack":     "\033[30m",
	"HiRed":       "\033[31m",
	"HiGreen":     "\033[32m",
	"HiYellow":    "\033[33m",
	"HiBlue":      "\033[34m",
	"HiPurple":    "\033[35m",
	"HiCyan":      "\033[36m",
	"HiWhite":     "\033[37m",
}

func colorize(msg string, vars interface{}) string {
	tpl, err := template.New(msg).Parse(msg)
	// Don't hold up the party for a bad template
	if err != nil {
		return msg
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return msg
	}

	return buf.String()
}
func Colorize(msg string) string {
	return colorize(msg, Colors)
}

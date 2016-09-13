// Inspiration (and some code) taken from https://github.com/deis/pkg/prettyprint
// The idea is not to rely on anything except the standard library

package logsip

import (
	"bytes"
	"text/template"
)

var Colors = map[string]string{
	"Default": "\033[0m",
	"Black":   "\033[0;30m",
	"Red":     "\033[0;31m",
	"Green":   "\033[0;32m",
	"Yellow":  "\033[0;33m",
	"Blue":    "\033[0;34m",
	"Purple":  "\033[0;35m",
	"Cyan":    "\033[0;36m",
	"White":   "\033[0;37m",
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

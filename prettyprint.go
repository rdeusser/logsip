package logsip

import (
	"bytes"
	"text/template"
)

var colors = map[string]string{
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

func prettyPrint(s string) string {
	tpl, err := template.New(s).Parse(s)
	// Don't hold up the party for a bad template
	if err != nil {
		return s
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, colors); err != nil {
		return s
	}

	return buf.String()
}

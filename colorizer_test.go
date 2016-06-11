package logsip

import "testing"

func TestColorize(t *testing.T) {
	out := Colorize("{{.Red}}Hello {{.Default}}World{{.UnderGreen}}!{{.Default}}")
	expected := "\033[0;31mHello \033[0mWorld\033[4;32m!\033[0m"
	if out != expected {
		t.Errorf("Expected '%s', got '%s'", expected, out)
	}
}

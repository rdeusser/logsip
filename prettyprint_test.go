package logsip

import "testing"

func TestPrettyPrint(t *testing.T) {
	out := prettyPrint("{{.Default}}The{{.Black}}quick{{.Red}}brown{{.Green}}fox{{.Yellow}}jumps{{.Blue}}over{{.Purple}}the{{.Cyan}}lazy{{.White}}dog{{.Default}}")
	expected := "\033[0mThe\033[0;30mquick\033[0;31mbrown\033[0;32mfox\033[0;33mjumps\033[0;34mover\033[0;35mthe\033[0;36mlazy\033[0;37mdog\033[0m"
	if out != expected {
		t.Errorf("Expected '%s', got '%s'", expected, out)
	}
}

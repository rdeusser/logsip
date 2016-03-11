package logsip_test

import (
	"bytes"
	"testing"

	"github.com/iamthemuffinman/logsip"
)

func TestPrint(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Print("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestPrintf(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Printf("%s", "foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestPrintln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Println("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfo(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Info("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfof(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Infof("%s", "foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfoln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Infoln("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarn(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Warn("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarnf(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Warnf("%s", "foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarnln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.New(&b)
	const testString = "foo"
	log.Warnln("foo")
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}

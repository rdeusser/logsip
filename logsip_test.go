package logsip_test

import (
	"bytes"
	"testing"

	"github.com/iamthemuffinman/logsip"
)

func TestPrint(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Print("foo")
	if expect := "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestPrintf(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Printf("%s", "foo")
	if expect := "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestPrintln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Println("foo")
	if expect := "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfo(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Info("foo")
	if expect := log.InfoPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfof(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Infof("%s", "foo")
	if expect := log.InfoPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestInfoln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Infoln("foo")
	if expect := log.InfoPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarn(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Warn("foo")
	if expect := log.WarnPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarnf(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Warnf("%s", "foo")
	if expect := log.WarnPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestWarnln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.Logger.SetOutput(&b)
	log.Warnln("foo")
	if expect := log.WarnPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestDebug(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.DebugMode = true
	log.Logger.SetOutput(&b)
	log.Debug("foo")
	if expect := log.DebugPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestDebugf(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.DebugMode = true
	log.Logger.SetOutput(&b)
	log.Debugf("%s", "foo")
	if expect := log.DebugPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestDebugln(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.DebugMode = true
	log.Logger.SetOutput(&b)
	log.Debugln("foo")
	if expect := log.DebugPrefix + "foo\n"; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}
func TestDebugMode(t *testing.T) {
	var b bytes.Buffer
	var log = logsip.Default()
	log.DebugMode = false
	log.Logger.SetOutput(&b)
	log.Debug("foo")
	log.Debugln("foo")
	log.Debugf("%s", "foo")
	if expect := ""; b.String() != expect {
		t.Errorf("There's an error here. %s should look like %s", expect, b.String())
	}
}

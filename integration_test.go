package logsip_test

import (
	"os"
	"os/exec"
	"testing"

	log "github.com/iamthemuffinman/logsip"
)

const testString = "hi"

func TestNew(t *testing.T) {
	logger := log.New()

	logger.Info(testString)
}

func TestDebug(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.Debug(testString)
}

func TestDebugf(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.Debugf("%s", testString)
}

func TestDebugln(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.Debugln(testString)
}

func TestInfo(t *testing.T) {
	log.Info(testString)
}

func TestInfof(t *testing.T) {
	log.Infof("%s", testString)
}

func TestInfoln(t *testing.T) {
	log.Infoln(testString)
}

func TestWarn(t *testing.T) {
	log.Warn(testString)
}

func TestWarnf(t *testing.T) {
	log.Warnf("%s", testString)
}

func TestWarnln(t *testing.T) {
	log.Warnln(testString)
}

func TestError(t *testing.T) {
	log.Error(testString)
}

func TestErrorf(t *testing.T) {
	log.Errorf("%s", testString)
}

func TestErrorln(t *testing.T) {
	log.Errorln(testString)
}

func TestFatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		log.Fatal(testString)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestFatalf(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		log.Fatal(testString)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatalf")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestFatalln(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		log.Fatal(testString)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatalln")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("it didn't panic")
		}
	}()

	log.Panic(testString)
}

func TestPanicf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("it didn't panic")
		}
	}()

	log.Panicf("%s", testString)
}

func TestPanicln(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("it didn't panic")
		}
	}()

	log.Panicln(testString)
}

func TestLevel(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.Debug(testString)
	l := log.GetLevel()
	if l != log.DebugLevel {
		t.Errorf("expected %s, got %s", log.DebugLevel, log.InfoLevel)
	}
	log.SetLevel(log.InfoLevel)
	log.Debug(testString)
	n := log.GetLevel()
	if n != log.InfoLevel {
		t.Errorf("expected %s, got %s", log.InfoLevel, log.DebugLevel)
	}
}

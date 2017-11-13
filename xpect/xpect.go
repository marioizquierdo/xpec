package xpect

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

func Xpect(t *testing.T, value interface{}) *Xpectation {
	return &Xpectation{
		value: value,
		t:     t,
	}
}

type Xpectation struct {
	t     *testing.T
	value interface{}
}

func (e *Xpectation) ToBe(expected interface{}) {
	// go1.9: e.t.Helper()
	if e.value != expected {
		e.failNow(fmt.Sprintf("expected %#v, found %#v", e.value, expected))
	}
}

func (e *Xpectation) failNow(msg string) {
	// go1.9: e.t.Helper()
	var lineText string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
		lineText, _, _ = readLine(file, line)
	} else {
		file = "???"
		lineText = "???"
		line = 1
	}

	e.t.Logf("%s:%d\n%s\n\t%s", file, line, lineText, msg)
	e.t.FailNow()
}

func readLine(fileName string, lineNum int) (line string, lastLine int, err error) {
	r, err := os.Open(fileName)
	if err != nil {
		return "", 0, err
	}

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

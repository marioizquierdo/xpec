package xpect

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func Xpect(t *testing.T, actual interface{}) *Xpectation {
	return &Xpectation{
		actual: actual,
		t:      t,
	}
}

type Xpectation struct {
	t      *testing.T
	actual interface{}
}

func (e *Xpectation) ToBe(expected interface{}) {
	// go1.9: e.t.Helper()
	if e.actual != expected {
		var msg string
		if reflect.TypeOf(e.actual) == reflect.TypeOf(expected) {
			msg = fmt.Sprintf("%#v is not %#v", e.actual, expected)
		} else {
			msg = fmt.Sprintf("%#v (%T) is not (%T) %#v", e.actual, e.actual, expected, expected)
		}
		e.failNow(msg)
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

	e.t.Logf("%s:%d\n%s\n%s", file, line, lineText, msg)
	e.t.FailNow()
}

func readLine(fileName string, lineNum int) (lineText string, lastLine int, err error) {
	r, err := os.Open(fileName)
	if err != nil {
		return "", 0, err
	}

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			lineText = strings.Trim(sc.Text(), " \t") // trim spaces and tabs
			return lineText, lastLine, sc.Err()
		}
	}
	return "", lastLine, io.EOF
}

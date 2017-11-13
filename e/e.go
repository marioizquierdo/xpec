package e

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

func Xpec(t *testing.T, value interface{}) *Subject {
	return &Subject{
		value: value,
		t:     t,
	}
}

type Subject struct {
	t     *testing.T
	value interface{}
}

func (s *Subject) ToBe(expected interface{}) {
	// go1.9: e.t.Helper()
	if s.value != expected {
		var msg string
		if reflect.TypeOf(s.value) == reflect.TypeOf(expected) {
			msg = fmt.Sprintf("but %#v is not %#v", s.value, expected)
		} else {
			msg = fmt.Sprintf("but (%T) %#v is not (%T) %#v", s.value, s.value, expected, expected)
		}
		s.failNow(msg)
	}
}

func (s *Subject) failNow(msg string) {
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

	s.t.Logf("%s:%d\n%s\n%s", file, line, lineText, msg)
	s.t.FailNow()
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

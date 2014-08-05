package log

import (
	. "github.com/tamaxyo/go-utils/testing"
	"testing"

	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

func TestLog(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0)
	l.Log(testString)
	file, line := linum() // line number matters
	expected := fmt.Sprintf("%s:%d: %s\n", filepath.Base(file), line-1, testString)
	EQUALS(t, "log should equal to", expected, b.String())
}

func TestLogWithMutipleArguments(t *testing.T) {
	const (
		arg1 = "foo"
		arg2 = "bar"
		arg3 = "baz"
	)

	var b bytes.Buffer
	l := New(&b, "", 0)
	l.Log(arg1, arg2, arg3)
	file, line := linum() // line number matters
	expected := fmt.Sprintf("%s:%d: %s %s %s\n", filepath.Base(file), line-1, arg1, arg2, arg3)
	EQUALS(t, "log should equal to", expected, b.String())
}

func TestCheckLogWrite(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0)

	err := errors.New("error")
	l.CheckLog(err, testString)
	file, line := linum() // line number matters
	expected := fmt.Sprintf("%s:%d: %s - %s\n", filepath.Base(file), line-1, testString, err)
	EQUALS(t, "log should equal to", expected, b.String())
}

func TestCheckLogWriteWithMultipleArguments(t *testing.T) {
	const (
		arg1 = "foo"
		arg2 = "bar"
		arg3 = "baz"
	)

	var b bytes.Buffer
	l := New(&b, "", 0)

	err := errors.New("error")
	l.CheckLog(err, arg1, arg2, arg3)
	file, line := linum() // line number matters
	expected := fmt.Sprintf("%s:%d: %s %s %s - %s\n", filepath.Base(file), line-1, arg1, arg2, arg3, err)
	EQUALS(t, "log should equal to", expected, b.String())
}

func TestCheckLogNoWrite(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0)

	l.CheckLog(nil, testString)
	EQUALS(t, "should equal", "", b.String())
}

//Standard functions can be called through Logger
func TestPrintln(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0)

	l.Println(testString)
	EQUALS(t, "should equal", testString+"\n", b.String())
}

func linum() (file string, line int) {
	_, file, line, _ = runtime.Caller(1)
	return
}

// Package testing adds helper functions for testing
package testing

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// ASSERT fails if condition is false
func ASSERT(tb testing.TB, condition bool, msg string) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("%s: %d: %s\n\n", filepath.Base(file), line, msg)
	}
}

// OK fails if err is not nil.
func OK(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
	}
}

// EQUALS fails if exp is not equal to act.
func EQUALS(tb testing.TB, msg string, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("%s:%d: %s\n\texp: %#v\n\tgot: %#v\n\n", filepath.Base(file), line, msg, exp, act)
	}
}

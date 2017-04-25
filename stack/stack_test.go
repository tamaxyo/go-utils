package stack_test

import (
	"testing"

	"github.com/tamaxyo/go-utils/stack"
	. "github.com/tamaxyo/go-utils/testing"
)

func TestLIFO(t *testing.T) {
	s := stack.NewStack()
	first := "first"
	second := "second"
	third := "third"

	s.Push(first)
	s.Push(second)
	s.Push(third)

	EQUALS(t, "string should match", third, s.Pop())
	EQUALS(t, "string should match", second, s.Pop())
	EQUALS(t, "string should match", first, s.Pop())
}

func TestPeek(t *testing.T) {
	s := stack.NewStack()
	first := "first"
	second := "second"
	third := "third"

	s.Push(first)
	s.Push(second)
	s.Push(third)

	EQUALS(t, "string should match", third, s.Peek())
	EQUALS(t, "string should match", third, s.Peek())

	EQUALS(t, "string should match", third, s.Pop())
	EQUALS(t, "string should match", second, s.Pop())
	EQUALS(t, "string should match", first, s.Pop())
}

func TestStackReturnsNilIfEmpty(t *testing.T) {
	s := stack.NewStack()

	EQUALS(t, "string should be empty", nil, s.Pop())
}

package calcrat

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/tamaxyo/go-utils/stack"
)

type priority int

const (
	high priority = iota
	low
)

// node is the interface that wraps val method.
type node interface {
	val() *big.Rat
}

// operator is the interface that groups basic functions of operator
type operator interface {
	setLeft(n node)
	setRight(n node)
	getPriority() priority
	cmp(op operator) int
}

var opMap map[string]func() operator = map[string]func() operator{
	"+": newAdd,
	"-": newSub,
	"*": newMul,
	"/": newDiv,
	"&": newBitwiseAnd,
	"|": newBitwiseOr,
	"^": newBitwiseXor,
}

type opBase struct {
	pri   priority
	left  node
	right node
}

// cmp returns the result of priority compiration
//  1 : higher
//  0 : same
// -1 : lower
func (a *opBase) cmp(b operator) int {
	if b == nil {
		return 1
	}

	if a.pri < b.getPriority() {
		return 1
	} else if a.pri > b.getPriority() {
		return -1
	}
	return 0
}

func (op *opBase) getPriority() priority {
	return op.pri
}

func (op *opBase) setLeft(n node) {
	op.left = n
}

func (op *opBase) setRight(n node) {
	op.right = n
}

type add struct {
	*opBase
}

func newAdd() operator {
	return &add{&opBase{low, nil, nil}}
}

func (op *add) val() *big.Rat {
	v := new(big.Rat)
	return v.Add(op.left.val(), op.right.val())
}

type sub struct {
	*opBase
}

func newSub() operator {
	return &sub{&opBase{low, nil, nil}}
}

func (op *sub) val() *big.Rat {
	v := new(big.Rat)
	return v.Sub(op.left.val(), op.right.val())
}

type mul struct {
	*opBase
}

func newMul() operator {
	return &mul{&opBase{high, nil, nil}}
}

func (op *mul) val() *big.Rat {
	v := new(big.Rat)
	return v.Mul(op.left.val(), op.right.val())
}

type div struct {
	*opBase
}

func newDiv() operator {
	return &div{&opBase{high, nil, nil}}
}

func (op *div) val() *big.Rat {
	v := new(big.Rat)
	inv := new(big.Rat)
	return v.Mul(op.left.val(), inv.Inv(op.right.val()))
}

// bitwiseAnd represents bitwise AND (&) operator.
// bitwiseAnd casts operands to uint64, so that incorrect value will be returned unless operands are uint64 compatible
type bitwiseAnd struct {
	*opBase
}

func newBitwiseAnd() operator {
	return &bitwiseAnd{&opBase{high, nil, nil}}
}

func (op *bitwiseAnd) val() *big.Rat {
	left := op.left.val().Num().Uint64() / op.left.val().Denom().Uint64()
	right := op.right.val().Num().Uint64() / op.right.val().Denom().Uint64()
	i := new(big.Int).SetUint64(left & right)
	return new(big.Rat).SetInt(i)
}

// bitwiseOr represents bitwise OR (|) operator.
// bitwiseOr casts operors to uint64, so that incorrect value will be returned unless operors are uint64 compatible
type bitwiseOr struct {
	*opBase
}

func newBitwiseOr() operator {
	return &bitwiseOr{&opBase{low, nil, nil}}
}

func (op *bitwiseOr) val() *big.Rat {
	left := op.left.val().Num().Uint64() / op.left.val().Denom().Uint64()
	right := op.right.val().Num().Uint64() / op.right.val().Denom().Uint64()
	i := new(big.Int).SetUint64(left | right)
	return new(big.Rat).SetInt(i)
}

// bitwiseXor represents bitwise XOR (^) operatxor.
// bitwiseXor casts operxors to uint64, so that incxorrect value will be returned unless operxors are uint64 compatible
type bitwiseXor struct {
	*opBase
}

func newBitwiseXor() operator {
	return &bitwiseXor{&opBase{low, nil, nil}}
}

func (op *bitwiseXor) val() *big.Rat {
	left := op.left.val().Num().Uint64() / op.left.val().Denom().Uint64()
	right := op.right.val().Num().Uint64() / op.right.val().Denom().Uint64()
	i := new(big.Int).SetUint64(left ^ right)
	return new(big.Rat).SetInt(i)
}

type literal struct {
	v *big.Rat
}

func newLiteral(s string, variables Variables, handler Handler) (*literal, error) {
	i := new(big.Int)
	l := &literal{
		v: new(big.Rat),
	}
	s = strings.TrimSpace(s)

	// Check if literal is int to accept hex and octal literals
	if _, ok := i.SetString(s, 0); ok {
		l.v.SetInt(i)
		return l, nil
	}

	if _, ok := l.v.SetString(s); ok {
		return l, nil
	}

	if v, ok := variables[s]; ok {
		l.v = v
		return l, nil
	}

	if handler != nil {
		if v := handler(s); v != nil {
			l.v = v
			return l, nil
		}
	}

	return nil, fmt.Errorf("could not parse string as rational - %s", s)
}

func (l *literal) val() *big.Rat {
	return l.v
}

// Calc returns the calculated rational value from given formula with given variables
func Calc(formula string, variables Variables, handler Handler) (*big.Rat, error) {
	re := regexp.MustCompile("\\(|\\)|\\+|-|\\*|/|&|\\||\\^|[^\\(\\)\\+\\-\\*/&\\|\\^]+")
	tokens := re.FindAllString(formula, -1)

	opStack := stack.NewStack()
	nodeStack := stack.NewStack()
	bracket := stack.NewStack()

	for _, token := range tokens {
		if fn, ok := opMap[token]; ok {
			op := fn()
			if f := opStack.Peek(); f != nil && op.cmp(f.(operator)) < 1 {
				prev := opStack.Pop().(operator)
				prev.setRight(nodeStack.Pop().(node))
				prev.setLeft(nodeStack.Pop().(node))
				nodeStack.Push(prev)
			}
			opStack.Push(op)
		} else if token == "(" {
			bracket.Push(opStack)
			opStack = stack.NewStack()
		} else if token == ")" {
			reduceBracket(opStack, nodeStack)
			opStack = bracket.Pop().(*stack.Stack)

		} else {
			l, err := newLiteral(token, variables, handler)
			if err == nil {
				nodeStack.Push(l)
			} else {
				return nil, fmt.Errorf("could not parse literal in the formula - formula: %s. detail: [%s]", formula, err.Error())
			}
		}
	}
	reduceBracket(opStack, nodeStack)

	return nodeStack.Pop().(node).val(), nil
}

func reduceBracket(opStack, nodeStack *stack.Stack) {
	for {
		p := opStack.Pop()
		if p == nil {
			return
		}
		op := p.(operator)
		op.setRight(nodeStack.Pop().(node))
		op.setLeft(nodeStack.Pop().(node))
		nodeStack.Push(op)
	}
}

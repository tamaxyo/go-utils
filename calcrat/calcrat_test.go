package calcrat_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/tamaxyo/go-utils/calcrat"
	. "github.com/tamaxyo/go-utils/testing"
)

func TestCalcCanEvaluateConstant(t *testing.T) {
	expected := big.NewRat(100, 1)
	actual, err := calcrat.Calc("100", nil, nil)
	EQUALS(t, "calc can evaluate constant", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateNamedValues(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error
	vars := map[string]*big.Rat{
		"one":   big.NewRat(100, 1),
		"two":   big.NewRat(200, 1),
		"three": big.NewRat(300, 1),
	}

	expected = big.NewRat(100, 1)
	actual, err = calcrat.Calc("one", vars, nil)
	EQUALS(t, "calc can evaluate first named value", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(200, 1)
	actual, err = calcrat.Calc("two", vars, nil)
	EQUALS(t, "calc can evaluate second named value", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateAddOperation(t *testing.T) {
	expected := big.NewRat(200, 1)
	actual, err := calcrat.Calc("100+100", nil, nil)
	EQUALS(t, "calc can evaluate add operation", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateSubOperation(t *testing.T) {
	expected := big.NewRat(50, 1)
	actual, err := calcrat.Calc("100-50", nil, nil)
	EQUALS(t, "calc can evaluate sub operation", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateMultOperation(t *testing.T) {
	expected := big.NewRat(10000, 1)
	actual, err := calcrat.Calc("100*100", nil, nil)
	EQUALS(t, "calc can evaluate mult operation", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateDivOperation(t *testing.T) {
	expected := big.NewRat(1, 1)
	actual, err := calcrat.Calc("100/100", nil, nil)
	EQUALS(t, "calc can evaluate div operation", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(1, 100)
	actual, err = calcrat.Calc("100/100/100", nil, nil)
	EQUALS(t, "calc can evaluate div operation", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateDivOperationWithFractionResult(t *testing.T) {
	expected := big.NewRat(3, 2)
	actual, err := calcrat.Calc("3/2", nil, nil)
	EQUALS(t, "calc can evaluate div operation with fraction result", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateAddOperationWhichContainsNamedValue(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	vars := map[string]*big.Rat{
		"one":   big.NewRat(100, 1),
		"two":   big.NewRat(200, 1),
		"three": big.NewRat(300, 1),
	}

	expected = big.NewRat(200, 1)
	actual, err = calcrat.Calc("100+one", vars, nil)
	EQUALS(t, "calc can evaluate add operation which contains first named value", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(300, 1)
	actual, err = calcrat.Calc("two+100", vars, nil)
	EQUALS(t, "calc can evaluate add operation which contains second named value", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(600, 1)
	actual, err = calcrat.Calc("one+two+three", vars, nil)
	EQUALS(t, "calc can evaluate add operation of named values", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateDivOperationWhichContainsNamedValue(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	vars := map[string]*big.Rat{
		"one":   big.NewRat(100, 1),
		"two":   big.NewRat(200, 1),
		"three": big.NewRat(300, 1),
	}

	expected = big.NewRat(1, 1)
	actual, err = calcrat.Calc("100/one", vars, nil)
	EQUALS(t, "calc can evaluate add operation which contains first named value", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(1, 2)
	actual, err = calcrat.Calc("two/400", vars, nil)
	EQUALS(t, "calc can evaluate add operation which contains second named value", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(1, 600)
	actual, err = calcrat.Calc("one/two/three", vars, nil)
	EQUALS(t, "calc can evaluate add operation of named values", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateOperationWithBrackets(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	expected = big.NewRat(20000, 1)
	actual, err = calcrat.Calc("(100+100)*100", nil, nil)
	EQUALS(t, "calc can evaluate operation with brackets ahead of mult", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(20000, 1)
	actual, err = calcrat.Calc("100*(100+100)", nil, nil)
	EQUALS(t, "calc can evaluate operation with brackets following to mult", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(2, 1)
	actual, err = calcrat.Calc("(100+100)/100", nil, nil)
	EQUALS(t, "calc can evaluate operation with brackets ahead of div", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(1, 2)
	actual, err = calcrat.Calc("100/(100+100)", nil, nil)
	EQUALS(t, "calc can evaluate operation with brackets following to div", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateOperationWithNestedBrackets(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	expected = big.NewRat(10000, 1)
	actual, err = calcrat.Calc("(100*(100-100)+100)*100", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested brackets with mult", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(10200, 1)
	actual, err = calcrat.Calc("100*((100+100)/100+100)", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested brackets with div", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(50, 1)
	actual, err = calcrat.Calc("((100+100)/100+8)/2*10", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested brackets with outer div and mult", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(100, 1)
	actual, err = calcrat.Calc("(10*(100+100))/10/2", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested brackets with outer divs", 0, expected.Cmp(actual))
	OK(t, err)

	expected = big.NewRat(21, 2)
	actual, err = calcrat.Calc("((10/(100+100))+1)*10", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested brackets with fraction result", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestAvoidZeroDivisionByAddingConstant(t *testing.T) {
	vars := map[string]*big.Rat{
		"one": big.NewRat(10, 1),
		"two": big.NewRat(0, 1),
	}

	expected := big.NewRat(10, 1)
	actual, err := calcrat.Calc("one/(two+1)", vars, nil)
	EQUALS(t, "avoid zero division by adding constant", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanEvaluateFormulaWithNestedFraction(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	expected = big.NewRat(2, 1)
	actual, err = calcrat.Calc("1/(5/10)", nil, nil)
	EQUALS(t, "calc can evaluate operation with nested fraction", expected.RatString(), actual.RatString())
	OK(t, err)
}

func TestCalcCanEvaluateFormulaWithNegativeResultInMid(t *testing.T) {
	var expected *big.Rat
	var actual *big.Rat
	var err error

	expected = big.NewRat(-89, 1)
	actual, err = calcrat.Calc("1+(10-100)", nil, nil)
	EQUALS(t, "calc can evaluate operation with negative result in mid", expected.RatString(), actual.RatString())
	OK(t, err)

	expected = big.NewRat(90*90, 1)
	actual, err = calcrat.Calc("(10-100)*(10-100)", nil, nil)
	EQUALS(t, "calc can evaluate operation with negative result in mid", expected.RatString(), actual.RatString())
	OK(t, err)
}

func TestCalcCanAcceptHexLiteral(t *testing.T) {
	expected := big.NewRat(255, 1)
	actual, err := calcrat.Calc("0xFF", nil, nil)
	EQUALS(t, "calc can evaluate hex literal", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestCalcCanAcceptOctalLiteral(t *testing.T) {
	expected := big.NewRat(63, 1)
	actual, err := calcrat.Calc("077", nil, nil)
	EQUALS(t, "calc can evaluate hex literal", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestReturnsErrorWithInvalidLiteral(t *testing.T) {
	_, err := calcrat.Calc("0xXX", nil, nil)
	ASSERT(t, "error should not be nil", err != nil)
}

func TestReturnsErrorWithUnknownVariable(t *testing.T) {
	var err error
	vars := map[string]*big.Rat{
		"one":   big.NewRat(100, 1),
		"two":   big.NewRat(200, 1),
		"three": big.NewRat(300, 1),
	}

	h := func(f string) *big.Rat {
		if f == "handle" {
			return big.NewRat(400, 1)
		}
		return nil
	}

	_, err = calcrat.Calc("four", vars, h)
	ASSERT(t, "error should not be nil", err != nil)

	_, err = calcrat.Calc("one+two+handle+four", vars, nil)
	ASSERT(t, "error should not be nil", err != nil)
}

func TestFormulaWithWhiteSpace(t *testing.T) {
	vars := map[string]*big.Rat{
		"one":   big.NewRat(100, 1),
		"two":   big.NewRat(200, 1),
		"three": big.NewRat(300, 1),
	}
	expected := big.NewRat(60100, 1)
	actual, err := calcrat.Calc("one + two   *   three", vars, nil)
	fmt.Printf("%#v\n", actual)
	EQUALS(t, "calc accepts formula with white space", 0, expected.Cmp(actual))
	OK(t, err)
}

func TestHandler(t *testing.T) {
	vars := map[string]*big.Rat{
		"one": big.NewRat(100, 1),
		"two": big.NewRat(200, 1),
	}
	expected := big.NewRat(60100, 1)
	h := func(f string) *big.Rat {
		if f == "three" {
			return big.NewRat(300, 1)
		}
		return nil
	}
	actual, err := calcrat.Calc("one + two   *   three", vars, h)
	fmt.Printf("%#v\n", actual)
	EQUALS(t, "calc accepts formula with white space", 0, expected.Cmp(actual))
	OK(t, err)
}

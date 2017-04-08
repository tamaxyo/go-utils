package sort_test

import (
	"fmt"
	"testing"

	"github.com/tamaxyo/go-utils/sort"
	. "github.com/tamaxyo/go-utils/testing"
)

func TestAscIntCanSortIntArrayInAscOrder(t *testing.T) {
	array := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	sort.IntAsc(array).Sort()

	EQUALS(t, "int array should be sorted in asc order", fmt.Sprintf("%#v", expected), fmt.Sprintf("%#v", array))
}

func TestAscIntCanSortIntArrayInDscOrder(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	sort.IntDsc(array).Sort()

	EQUALS(t, "int array should be sorted in dsc order", fmt.Sprintf("%#v", expected), fmt.Sprintf("%#v", array))
}

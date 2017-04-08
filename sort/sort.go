package sort

import (
	"sort"
)

type IntAsc []int
type IntDsc []int

func (asc IntAsc) Sort()              { sort.Sort(asc) }
func (asc IntAsc) Len() int           { return len(asc) }
func (asc IntAsc) Swap(i, j int)      { asc[i], asc[j] = asc[j], asc[i] }
func (asc IntAsc) Less(i, j int) bool { return asc[i] < asc[j] }

func (dsc IntDsc) Sort()              { sort.Sort(dsc) }
func (dsc IntDsc) Len() int           { return len(dsc) }
func (dsc IntDsc) Swap(i, j int)      { dsc[i], dsc[j] = dsc[j], dsc[i] }
func (dsc IntDsc) Less(i, j int) bool { return dsc[i] > dsc[j] }

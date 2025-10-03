package sorter

import (
	"fmt"
	"sort"
	"wbsort/internal/parser"
)

type Sorter struct {
	flags *parser.Flags
	Lines []string
}

func New(flags *parser.Flags, Lines []string) *Sorter {
	return &Sorter{
		flags: flags,
		Lines: Lines,
	}
}

func (s *Sorter) Sort() {
	if s.flags.FlagC {
		if ok := s.isSorted(); ok {
			fmt.Println("values are sorted")
			return
		}
		fmt.Println("values are not sorted")
		return
	}
	if s.flags.FlagB {
		s.ignoreTrailingBlanks()
	}

	if s.flags.FlagU {
		s.setUnique()
	}

	s.baseSort()

	if s.flags.FlagK > 0 {
		s.sortByColumn()
	}

	if s.flags.FlagN {
		s.sortByNumeric()
	}

	if s.flags.FlagM {
		s.sortByMonth()
	}

	if s.flags.FlagH {
		s.sortByHumanSuffix()
	}

	if s.flags.FlagR {
		s.sortReverse()
	}
}

func (s *Sorter) baseSort() {
	sort.Slice(s.Lines, func(i, j int) bool {
		return s.Lines[i] < s.Lines[j]
	})
}

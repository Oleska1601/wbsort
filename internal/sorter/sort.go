package sorter

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/Oleska1601/wbsort/internal/parser"
)

// Sorter implements sorting functionality with various algorithms
type Sorter struct {
	flags *parser.Flags
	lines []string
}

// New creates and returns a new Sorter instance
func New(flags *parser.Flags, lines []string) *Sorter {
	return &Sorter{
		flags: flags,
		lines: lines,
	}
}

// Sort performs sorting according to configured flags
func (s *Sorter) Sort() []string {
	if s.flags.FlagU {
		s.SetUnique()
	}

	if s.flags.FlagC {
		if !s.IsSorted() {
			return []string{"values are not sorted"}
		}

		return nil
	}

	sort.SliceStable(s.lines, s.Less)

	return s.lines
}

// IsSorted checks if lines are already in sorted order
func (s *Sorter) IsSorted() bool {
	for i := 1; i < len(s.lines); i++ {
		if !s.Less(i-1, i) {
			fmt.Println(i-1, " ", i)
			return false
		}
	}
	return true
}

// Less reports whether element i should sort before element j
func (s *Sorter) Less(i, j int) bool {
	vi := s.lines[i]
	vj := s.lines[j]

	if s.flags.FlagB {
		vi = strings.TrimRightFunc(vi, unicode.IsSpace)
		vj = strings.TrimRightFunc(vj, unicode.IsSpace)
	}

	vi = getSortKey(vi, s.flags.FlagK)
	vj = getSortKey(vj, s.flags.FlagK)

	// Выбираем стратегию сравнения
	var less bool

	switch {
	case s.flags.FlagM:
		less = compareByMonth(vi, vj)
	case s.flags.FlagH:
		less = compareByHumanSuffix(vi, vj)
	case s.flags.FlagN:

		less = compareByNumeric(vi, vj)
	default:
		// base sort
		less = vi < vj
	}

	if s.flags.FlagR {
		return !less
	}

	return less

}

// SetUnique removes duplicate lines from the dataset
func (s *Sorter) SetUnique() {
	result := make([]string, 0, len(s.lines))
	m := make(map[string]struct{})
	for _, line := range s.lines {
		key := line
		if s.flags.FlagK > 0 {
			key = getSortKey(line, s.flags.FlagK)
		}

		if _, ok := m[key]; !ok {
			result = append(result, line)
			m[key] = struct{}{}
		}
	}

	s.lines = result
}

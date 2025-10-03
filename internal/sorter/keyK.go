package sorter

import (
	"sort"
	"strings"
)

func (s *Sorter) sortByColumn() {
	if len(s.Lines) == 0 {
		return
	}

	n := s.flags.FlagK
	if n > 0 {
		n -= 1
	}

	sort.Slice(s.Lines, func(i, j int) bool {
		return compareByColumn(s.Lines[i], s.Lines[j], n)
	})
}

func compareByColumn(firstLine, secondLine string, n int) bool {

	firstFields := strings.Fields(firstLine)
	secondFields := strings.Fields(secondLine)

	var firstVal, secondVal string
	if n < len(firstFields) {
		firstVal = firstFields[n]
	}
	if n < len(secondFields) {
		secondVal = secondFields[n]
	}
	return firstVal < secondVal

}

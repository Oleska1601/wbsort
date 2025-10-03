package sorter

import (
	"sort"
	"strconv"
	"strings"
)

func (s *Sorter) sortByNumeric() {
	if len(s.Lines) == 0 {
		return
	}
	n := s.flags.FlagK
	if n > 0 {
		n -= 1
	}
	sort.Slice(s.Lines, func(i, j int) bool {
		return compareByNumeric(s.Lines[i], s.Lines[j], n)
	})
}

func compareByNumeric(firstLine, secondLine string, n int) bool {
	firstFields := strings.Fields(firstLine)
	secondFields := strings.Fields(secondLine)

	var firstVal, secondVal string
	if n < len(firstFields) {
		firstVal = firstFields[n]
	}
	if n < len(secondFields) {
		secondVal = secondFields[n]
	}
	firstNum, err1 := strconv.Atoi(firstVal)
	secondNum, err2 := strconv.Atoi(secondVal)

	switch {
	case err1 == nil && err2 == nil:
		return firstNum < secondNum
	case err1 == nil:
		return true
	case err2 == nil:
		return false
	default:
		return firstVal < secondVal
	}

}

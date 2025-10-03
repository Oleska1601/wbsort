package sorter

import (
	"sort"
	"strings"
)

func (s *Sorter) sortByMonth() {
	if len(s.Lines) == 0 {
		return
	}
	n := s.flags.FlagK
	if n > 0 {
		n -= 1
	}
	sort.Slice(s.Lines, func(i, j int) bool {
		return s.compareByMonth(s.Lines[i], s.Lines[j], n)
	})

}

func (s *Sorter) compareByMonth(firstLine, secondLine string, n int) bool {
	firstFields := strings.Fields(firstLine)
	secondFields := strings.Fields(secondLine)

	var firstVal, secondVal string
	if n < len(firstFields) {
		firstVal = firstFields[n]
	}
	if n < len(secondFields) {
		secondVal = secondFields[n]
	}
	firstMonth, ok1 := parseMonth(firstVal)
	secondMonth, ok2 := parseMonth(secondVal)

	switch {
	case ok1 && ok2:
		return firstMonth < secondMonth
	case ok1:
		return true
	case ok2:
		return false
	default:
		return firstVal < secondVal
	}

}

func parseMonth(inputLine string) (int, bool) {
	if len(inputLine) == 0 {
		return 0, false
	}
	line := strings.ToLower(strings.TrimSpace(inputLine))
	words := strings.Fields(line)
	if len(line) >= 3 {
		if month, exists := MonthMapping[line[:3]]; exists {
			return month, true
		}
	}

	for _, word := range words {
		if len(word) >= 3 {
			if monthIdx, exists := MonthMapping[word[:3]]; exists {
				return monthIdx, true
			}
		}
	}

	return 0, false
}

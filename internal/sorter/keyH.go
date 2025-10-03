package sorter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (s *Sorter) sortByHumanSuffix() {
	if len(s.Lines) == 0 {
		return
	}
	n := s.flags.FlagK
	if n > 0 {
		n -= 1
	}
	sort.Slice(s.Lines, func(i, j int) bool {
		return compareByHumanSuffix(s.Lines[i], s.Lines[j], n)
	})
}

func compareByHumanSuffix(firstLine, secondLine string, n int) bool {
	firstFields := strings.Fields(firstLine)
	secondFields := strings.Fields(secondLine)
	var firstVal, secondVal string
	if n < len(firstFields) {
		firstVal = firstFields[n]
	}
	if n < len(secondFields) {
		secondVal = secondFields[n]
	}
	val1, err1 := parseHumanVal(firstVal)
	val2, err2 := parseHumanVal(secondVal)
	switch {
	case err1 == nil && err2 == nil:
		return val1 < val2
	case err1 == nil:
		return true
	case err2 == nil:
		return false
	default:
		return firstVal < secondVal
	}
}

func parseHumanVal(inputLine string) (float64, error) {
	if len(inputLine) == 0 {
		return 0, fmt.Errorf("len(inputLine) == 0")
	}
	line := strings.TrimSpace(strings.ToLower(inputLine))
	lastVal := inputLine[len(inputLine)-1]
	if multiplier, ok := MultiplierMapping[lastVal]; ok {
		numStr := strings.TrimSpace(inputLine[:len(inputLine)-1])
		numVal, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, fmt.Errorf("parse float error: %w", err)
		}
		result := numVal * multiplier
		return result, nil
	}
	// если нет суффикса -> просто обычное число
	numVal, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float error: %w", err)
	}
	return numVal, nil
}

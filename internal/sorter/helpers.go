package sorter

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var months = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

var human = map[string]float64{
	"k": 1024,
	"m": 1024 * 1024,
	"g": 1024 * 1024 * 1024,
	"t": 1024 * 1024 * 1024 * 1024,
}

func getSortKey(s string, col int) string {
	if col <= 0 {
		return s
	}

	fields := strings.Split(s, "\t")
	if col-1 < len(fields) {
		return fields[col-1]
	}

	return ""
}

// -n
func compareByNumeric(firstVal, secondVal string) bool {
	re := regexp.MustCompile(`(-?\d+\.?\d*)|([^\d]+)`)
	parts1 := re.FindAllString(firstVal, -1)
	parts2 := re.FindAllString(secondVal, -1)

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		firstNum, err1 := strconv.ParseFloat(parts1[i], 64)
		secondNum, err2 := strconv.ParseFloat(parts2[i], 64)
		switch {
		case err1 == nil && err2 == nil:
			return firstNum < secondNum
		default:
			continue
		}
	}

	return firstVal < secondVal
}

// -h
func parseHumanVal(str string) (float64, error) {
	re := regexp.MustCompile(`^(\d+(?:\.\d+)?)(k|m|g|t)?$`)
	// find string submatch - найти 1 совпадение с группами: 1K = [1k, 1, k]
	match := re.FindStringSubmatch(strings.ToLower(str))

	if match == nil {
		return 0, errors.New("invalid format")
	}

	numStr := match[1]
	suffix := match[2]

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		panic(err)
	}

	switch suffix {
	case "":
		return num, nil
	default:
		multiplier, ok := human[suffix]
		if !ok {
			return 0, errors.New("invalid format")
		}
		return num * multiplier, nil
	}
}

func compareByHumanSuffix(firstVal, secondVal string) bool {
	val1, err1 := parseHumanVal(firstVal)
	val2, err2 := parseHumanVal(secondVal)
	switch {
	case err1 == nil && err2 == nil:
		return val1 < val2
	case err1 == nil && err2 != nil:
		return false
	case err2 == nil && err1 != nil:
		return true
	default:
		return firstVal < secondVal
	}
}

// -M
func parseMonth(val string) (int, bool) {
	line := strings.ToLower(val)

	// Ищем любое 3-буквенное сочетание, которое является месяцем
	for i := 0; i <= len(line)-3; i++ {
		substr := line[i : i+3]
		if month, exists := months[substr]; exists {
			return month, true
		}
	}

	return 0, false
}

func compareByMonth(firstVal, secondVal string) bool {
	firstMonth, ok1 := parseMonth(firstVal)
	secondMonth, ok2 := parseMonth(secondVal)

	switch {
	case ok1 && ok2:
		return firstMonth < secondMonth
	case ok1 && !ok2:
		return true
	case ok2 && !ok1:
		return false
	default:
		return firstVal < secondVal
	}

}

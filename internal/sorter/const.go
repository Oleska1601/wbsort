package sorter

var MonthMapping = map[string]int{
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

var MultiplierMapping = map[byte]float64{
	'k': 1024,
	'm': 1024 * 1024,
	'g': 1024 * 1024 * 1024,
	't': 1024 * 1024 * 1024 * 1024,
}

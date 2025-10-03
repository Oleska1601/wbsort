package sorter

import(
	"testing"
	"wbsort/internal/parser"
)

func TestSorter_ignoreTrailingBlanks(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags *parser.Flags
		Lines []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.flags, tt.Lines)
			s.ignoreTrailingBlanks()
		})
	}
}

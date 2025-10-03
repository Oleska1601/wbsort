package sorter

import (
	"strings"
	"unicode"
)

func (s *Sorter) ignoreTrailingBlanks() {
	for i := 0; i < len(s.Lines); i++ {
		s.Lines[i] = strings.TrimRightFunc(s.Lines[i], unicode.IsSpace)
	}
}

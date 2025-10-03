package sorter

import "slices"

func (s *Sorter) sortReverse() {
	slices.Reverse(s.Lines)
}

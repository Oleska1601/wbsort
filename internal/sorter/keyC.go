package sorter

func (s *Sorter) isSorted() bool {
	for i := 1; i < len(s.Lines); i++ {
		if s.Lines[i-1] > s.Lines[i] {
			return false
		}
	}
	return true
}

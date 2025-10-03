package sorter

func (s *Sorter) setUnique() {
	s.Lines = removeDuplicates(s.Lines)
}

func removeDuplicates(sl []string) []string {
	var result []string
	m := make(map[string]struct{})
	for _, s := range sl {
		if _, ok := m[s]; !ok {
			result = append(result, s)
			m[s] = struct{}{}
		}
	}
	return result
}

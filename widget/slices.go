package widget

func sliceDeduplicate[S ~[]E, E comparable](s S) []E {
	seen := make(map[E]bool)
	s2 := make([]E, 0)
	for _, v := range s {
		if seen[v] {
			continue
		}
		s2 = append(s2, v)
		seen[v] = true
	}
	return s2
}

func sliceDeleteFunc[S ~[]E, E comparable](s S, delete func(v E) bool) []E {
	s2 := make([]E, 0)
	for _, x := range s {
		if delete(x) {
			continue
		}
		s2 = append(s2, x)
	}
	return s2
}

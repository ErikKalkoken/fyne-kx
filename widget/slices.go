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

// sliceDeleteFunc is a re-implementation of slices.DeleteFunc for Go 1.19.
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

// sliceClone is a re-implementation of slices.Clone for Go 1.19.
func sliceClone[S ~[]E, E comparable](s S) []E {
	s2 := make([]E, len(s))
	copy(s2, s)
	return s2
}

// sliceContains is a re-implementation of slices.Contains for Go 1.19.
func sliceContains[S ~[]E, E comparable](s S, v E) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

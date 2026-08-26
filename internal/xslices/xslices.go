// Package xslices contains helper functions for slices.
package xslices

import "slices"

// Deduplicate returns a new slice where all duplicate elements have been removed.
// The order of the elements is not changed, but the new slice can be shorter.
func Deduplicate[S ~[]E, E comparable](s S) []E {
	seen := make(map[E]bool, len(s))
	s2 := make([]E, 0, len(s))
	for _, v := range s {
		if seen[v] {
			continue
		}
		s2 = append(s2, v)
		seen[v] = true
	}
	return slices.Clip(s2)
}

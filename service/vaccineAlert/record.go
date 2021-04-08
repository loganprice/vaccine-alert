package vaccineAlert

import "github.com/loganprice/vaccine-alert/pkg/types"

func Difference(current, previous []types.AvailableSite) []types.AvailableSite {
	var diff []types.AvailableSite

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range current {
			found := false
			for _, s2 := range previous {
				if s1.Name == s2.Name {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			current, previous = previous, current
		}
	}
	return diff
}

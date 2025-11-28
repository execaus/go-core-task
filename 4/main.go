package main

// DifferenceStrings возвращает слайс строк, содержащий элементы slice1, которых нет в slice2.
func DifferenceStrings(slice1, slice2 []string) []string {
	diff := make([]string, 0)
	m := make(map[string]bool)

	for _, s := range slice2 {
		m[s] = true
	}

	for _, s := range slice1 {
		if !m[s] {
			diff = append(diff, s)
		}
	}

	return diff
}

package main

// IntersectInts проверяет, есть ли пересечение между двумя срезами целых чисел.
func IntersectInts(a, b []int) (bool, []int) {
	setA := make(map[int]struct{})
	for _, val := range a {
		setA[val] = struct{}{}
	}

	intersectionSet := make(map[int]struct{})
	for _, val := range b {
		if _, exists := setA[val]; exists {
			intersectionSet[val] = struct{}{}
		}
	}

	if len(intersectionSet) == 0 {
		return false, nil
	}

	intersection := make([]int, 0, len(intersectionSet))
	for val := range intersectionSet {
		intersection = append(intersection, val)
	}

	return true, intersection
}

package utils

import "fmt"

func SliceIndexOf[T any](
	s []T,
	fn func(T) bool,
) (int, error) {
	for i, el := range s {
		if fn(el) {
			return i, nil
		}
	}

	return -1, fmt.Errorf("failed to find a matching element")
}

func SliceContains[T comparable](s []T, el T) bool {
	for _, element := range s {
		if element == el {
			return true
		}
	}

	return false
}

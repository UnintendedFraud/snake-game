package utils

import "fmt"

func IndexOf[T any](
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

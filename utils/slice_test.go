package utils

import "testing"

type complexStruct struct {
	field string
	value int
}

func Test_IndexOf(t *testing.T) {
	type expected struct {
		idx       int
		errIsNull bool
	}
	type example[T any] struct {
		name     string
		s        []T
		searchFn func(T) bool
		expected expected
	}

	strSlice := []string{"one", "two", "three"}
	complexSlice := []complexStruct{
		{field: "test", value: 12},
		{field: "something", value: 333},
		{field: "else", value: 3432},
		{field: "foo", value: 2},
	}

	examples := []example[string]{
		{
			name:     "valid index in slice of strings",
			s:        strSlice,
			searchFn: func(s string) bool { return s == "three" },
			expected: expected{idx: 2, errIsNull: true},
		},
		{
			name:     "invalid index in slice of strings",
			s:        strSlice,
			searchFn: func(s string) bool { return s == "invalid" },
			expected: expected{idx: -1, errIsNull: false},
		},
	}

	cplxExamples := []example[complexStruct]{
		{
			name:     "valid index in slice of custom struct",
			s:        complexSlice,
			searchFn: func(el complexStruct) bool { return el.value == 333 },
			expected: expected{idx: 1, errIsNull: true},
		},
		{
			name:     "invalid index in slice of custom struct",
			s:        complexSlice,
			searchFn: func(el complexStruct) bool { return el.field == "invalid" },
			expected: expected{idx: -1, errIsNull: false},
		},
	}

	for _, ex := range examples {
		idx, err := SliceIndexOf(ex.s, ex.searchFn)

		if idx != ex.expected.idx {
			t.Errorf("%s: expected index [%d] but got [%d]", ex.name, ex.expected.idx, idx)
		}

		errIsNull := err == nil
		if errIsNull != ex.expected.errIsNull {
			t.Errorf("%s: expected nil error [%t] but got [%t]", ex.name, ex.expected.errIsNull, errIsNull)
		}
	}

	for _, ex := range cplxExamples {
		idx, err := SliceIndexOf(ex.s, ex.searchFn)

		if idx != ex.expected.idx {
			t.Errorf("%s: expected index [%d] but got [%d]", ex.name, ex.expected.idx, idx)
		}

		errIsNull := err == nil
		if errIsNull != ex.expected.errIsNull {
			t.Errorf("%s: expected nil error [%t] but got [%t]", ex.name, ex.expected.errIsNull, errIsNull)
		}
	}
}

func Test_Contains(t *testing.T) {
	type examples[T comparable] struct {
		s        []T
		values   []T
		expected []bool
	}

	examples1 := examples[int]{
		s:        []int{2, 3, 432, 123, 21},
		values:   []int{432, 111},
		expected: []bool{true, false},
	}
	examples2 := examples[string]{
		s:        []string{"one", "two", "three"},
		values:   []string{"invalid", "one"},
		expected: []bool{false, true},
	}

	for i, v := range examples1.values {
		expected := examples1.expected[i]

		if SliceContains(examples1.s, v) != expected {
			t.Errorf("Test_Contains - [%v] being in %+v should be %t", v, examples1.s, expected)
		}
	}

	for i, v := range examples2.values {
		expected := examples2.expected[i]

		if SliceContains(examples2.s, v) != expected {
			t.Errorf("Test_Contains - [%v] being in %+v should be %t", v, examples2.s, expected)
		}
	}
}

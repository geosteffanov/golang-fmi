package main

import "strings"

func Repeater(s, sep string) func(int) string {
	const EmptyString = ""

	return func(times int) string {
		if times < 0 {
			return EmptyString
		}

		wordRepetitions := make([]string, times)

		for i := 0; i < times; i++ {
			wordRepetitions[i] = s
		}

		return strings.Join(wordRepetitions, sep)
	}
}

func Generator(gen func(int) int, initial int) func() int {
	current := initial
	firstInvocation := true

	return func() int {
		if firstInvocation {
			firstInvocation = false
			return current
		}

		current = gen(current)
		return current
	}
}

func MapReducer(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {
	return func(ints ...int) int {
		result := initial

		for _, int := range ints {
			result = reducer(result, mapper(int))
		}

		return result
	}

}

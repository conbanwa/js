package js

import "fmt"

// Slice
//
// start Optional
// Zero-based index at which to start extraction, converted to an integer.
//
// Negative index counts back from the end of the array — if start < 0, start + array.length is used.
// If start < -array.length or start is omitted, 0 is used.
// If start >= array.length, nothing is extracted.
//
// end Optional
// Zero-based index at which to end extraction, converted to an integer. slice() extracts up to but not including end.
//
// Negative index counts back from the end of the array — if end < 0, end + array.length is used.
// If end < -array.length, 0 is used.
// If end >= array.length or end is omitted, array.length is used, causing all elements until the end to be extracted.
// If end is positioned before or at start after normalization, nothing is extracted.
func Slice[T comparable](arr []T, limit ...int) []T {
	start := 0
	end := len(arr)
	if len(limit) > 0 {
		start = limit[0]
		if len(limit) > 1 {
			end = limit[1]
		}
	}
	switch true {
	case start <= -len(arr):
		start = 0
	case start < 0:
		start = len(arr) + start
	case start >= len(arr):
		return arr[0:0]
	}
	switch true {
	case end <= -len(arr):
		end = 0
	case end < 0:
		end = len(arr) + end
	case end > len(arr):
		end = len(arr)
	}
	if start > end {
		return arr[0:0]
	}
	return arr[start:end]
}

// SubString
// similar with Slice
func SubString(arr string, limit ...int) string {
	start := 0
	end := len(arr)
	if len(limit) > 0 {
		start = limit[0]
		if len(limit) > 1 {
			end = limit[1]
		}
	}
	switch true {
	case start <= -len(arr):
		start = 0
	case start < 0:
		start = len(arr) + start
	case start >= len(arr):
		return arr[0:0]
	}
	switch true {
	case end <= -len(arr):
		end = 0
	case end < 0:
		end = len(arr) + end
	case end > len(arr):
		end = len(arr)
	}
	if start > end {
		return arr[0:0]
	}
	return arr[start:end]
}

func Table[T any](arr []T, args ...any) string {
	out := fmt.Sprint(args...) + "\n"
	for i, a := range arr {
		out += fmt.Sprintf("%d|%+v\n", i, a)
	}
	return out
}

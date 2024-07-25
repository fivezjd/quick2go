package syntaxBase

import "golang.org/x/exp/constraints"

func CompareGetMax[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

package gofunctional

import "fmt"

func Map[A, B any](f func(A) B, xs []A) []B {
	ys := make([]B, len(xs))

	for i, v := range xs {
		ys[i] = f(v)
	}

	return ys
}

func Reduce[A any](f func(A, A) A, xs []A) (A, error) {
	var result A
	if len(xs) == 0 {
		return result, fmt.Errorf("slice cannot be empty")
	}
	result = xs[0]

	for _, v := range xs[1:] {
		result = f(result, v)
	}

	return result, nil
}

func Filter[A any](f func(A) bool, xs []A) []A {
	ys := make([]A, 0, len(xs))

	for _, v := range xs {
		if f(v) {
			ys = append(ys, v)
		}
	}

	return ys
}

package streams

func Map[A, B any](f func(A) B, in <-chan A) <-chan B {
	out := make(chan B)

	go func() {
		defer close(out)
		for elem := range in {
			out <- f(elem)
		}
	}()

	return out
}

func Filter[A any](pred func(A) bool, in <-chan A) <-chan A {
	out := make(chan A)

	go func() {
		defer close(out)
		for elem := range in {
			if pred(elem) {
				out <- elem
			}
		}
	}()

	return out
}

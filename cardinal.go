package aviary

// Cardinal :: (a -> b -> c) -> b -> a -> c
//
// C combinator or flip
func Cardinal[T, C any](f func(T) func(T) C) func(T) func(T) C {
	return func(b T) func(T) C {
		return func(a T) C {
			return f(b)(a)
		}
	}
}
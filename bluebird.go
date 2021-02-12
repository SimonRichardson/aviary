package aviary

// Bluebird :: (b -> c) -> (a -> b) -> a -> c
//
// B combinator or function composition
func Bluebird[A, B, C any](f func(B) C) func(func(A) B) func(A) C {
	return func(g func(A) B) func(A) C {
		return func(x A) C {
			return f(g(x))
		}
	}
}

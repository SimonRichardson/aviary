package aviary

// Becard :: (c -> d) -> (b -> c) -> (a -> b) -> a -> d
//
// B3 combinator or function composition (for three functions)
func Becard[A, B, C, D any](f func(C) D) func(func(B) C) func(func(A) B) func(A) D {
	return func(g func(B) C) func(func(A) B) func(A) D {
		return func(h func(A) B) func(A) D {
			return func(x A) D {
				return f(g(h(x)))
			}
		}
	}
}

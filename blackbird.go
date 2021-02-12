package aviary

// Blackbird :: (c -> d) -> (a -> b -> c) -> a -> b -> d
//
// B1 combinator
func Blackbird[A, B, C, D any](f func(C) D) func(func(A, B) C) func(A, B) D {
	return func(g func(A, B) C) func(A, B) D {
		return func(x A, y B) D {
			return f(g(x, y))
		}
	}
}

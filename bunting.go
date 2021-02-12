package aviary

// Bunting :: (d -> e) -> (a -> b -> c -> d) -> a -> b -> c -> e
//
// B2 combinator
func Bunting[A, B, C, D, E any](f func(D) E) func(func(A, B, C) D) func(A) func(B) func(C) E {
	return func(g func(A, B, C) D) func(A) func(B) func(C) E {
		return func(x A) func(B) func(C) E {
			return func(y B) func(C) E {
				return func(z C) E {
					return f(g(x, y, z))
				}
			}
		}
	}
}

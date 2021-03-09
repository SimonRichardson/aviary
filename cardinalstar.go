package aviary

// CardinalStar :: (a -> c -> b -> d) -> a -> b -> c -> d
//
// C* combinator - cardinal once removed.
func CardinalStar[T, D any](f func(T) func(T) func(T) D) func(T) func(T) func(T) D {
	return func(x T) func(T) func(T) D {
		return func(y T) func(T) D {
			return func(z T) D {
				return f(x)(z)(y)
			}
		}
	}
}